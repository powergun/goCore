package reGex

import (
	"errors"
	"strings"
	"testing"
	"github.com/powergun/goUsage/_resources"
	"os"
	"bufio"
)

func parseAndAssert(t *testing.T, line string, field string, expected string) {
	if m, _, err := Parse(line); err != nil {
		t.Fatalf("failed to parse: %s", line)
	} else {
		if v, OK := m[field]; !OK {
			t.Fatalf("%s field  does not exist", field)
		} else {
			if v != expected {
				t.Fatalf("expecting %s, got: %s", expected, v)
			}
		}
	}
}

func iterLines(fileName string, callback func(string, int)error) error {
	filePath := _resources.ResourcePath(fileName)
	f, _ := os.Open(filePath)
	scanner := bufio.NewScanner(f)
	idx := 1
	for scanner.Scan() {
		err := callback(scanner.Text(), idx)
		if err != nil {
			return err
		}
		idx++
	}
	return nil
}

func TestEmptyLine(t *testing.T) {
	assertEmptyLine := func(line string) {
		if _, isEmptyLine, _ := Parse(line); isEmptyLine == false {
			t.Fatalf("fail to identify empty line")
		}
	}
	assertEmptyLine("")
	assertEmptyLine("\n")
	assertEmptyLine("  ")
	assertEmptyLine("     \n")
}

func TestHasNoSubroutineInfo(t *testing.T) {
	parseAndAssert(t, "#2 <signal handler called>", "Id", "2")
}

func TestCaptureId(t *testing.T) {
	parseAndAssert(t, "#0 ?? ()", "Id", "0")
	parseAndAssert(t, " asd #0 23 ?? ()", "Id", "0")
	parseAndAssert(t, "#0  m4_traceon (obs=0x24e)", "Id", "0")
	parseAndAssert(t, "#1  0x6e38 in expand_macro (sym=0x2b600) at macro.c:242", "Id", "1")
	parseAndAssert(t, "#2  0x6840 in expand_token (obs=0x0, t=177664, td=0xf7fffb08)", "Id", "2")
}

func TestCaptureAddress(t *testing.T) {
	parseAndAssert(t, "#0 ?? ()", "Address", "")
	parseAndAssert(t, "#0  m4_traceon (obs=0x24e)", "Address", "")
	parseAndAssert(t, "#1  0x6e38 in expand_macro (sym=0x2b) at macro.c:242", "Address", "0x6e38")
	parseAndAssert(t, "#2  0x6840 in expand_token (obs=0x0, t=17, td=0xf7fffb08)", "Address", "0x6840")
}

func TestCaptureSubroutine(t *testing.T) {
	parseAndAssert(t, "#0 ?? ()", "Subroutine", "??")
	parseAndAssert(t, "#0  m4_traceon (obs=0x24e)", "Subroutine", "m4_traceon")
	parseAndAssert(t, "#1  0x68 in expand_macro (sym=0x2b) at mo.c:242", "Subroutine", "expand_macro")
	parseAndAssert(t, "#2  0x40 in expand_token (obs=0x0, t=17, tdxb08)", "Subroutine", "expand_token")
}

func TestCaptureSubroutineWithNamespace(t *testing.T) {
	parseAndAssert(t, "#0 SomeModule::SomeClass::someMeth ()",
		"Subroutine", "SomeModule::SomeClass::someMeth")
}

func TestCaptureSubroutineWithCTOR(t *testing.T) {
	parseAndAssert(t, "#4  0x00007fdef42ddbcd in IlmThread_2_2::Semaphore::wait() () from /usr/lib/x86_64-linux-gnu/libIlmThread-2_2.so.12",
		"Subroutine", "IlmThread_2_2::Semaphore::wait()")

	// excessive space:
	// expect subroutine and arguments are handled properly
	parseAndAssert(t, "#4  0x00007fdef42ddbcd in IlmThread_2_2::Semaphore::wait () () from /usr/lib/x86_64-linux-gnu/libIlmThread-2_2.so.12",
		"Subroutine", "IlmThread_2_2::Semaphore::wait")
	parseAndAssert(t, "#4  0x00007fdef42ddbcd in IlmThread_2_2::Semaphore::wait () () from /usr/lib/x86_64-linux-gnu/libIlmThread-2_2.so.12",
		"Arguments", "()")
}

func TestCaptureSubroutineWithCTORParameters(t *testing.T) {
	parseAndAssert(t, "#0 SomeModule::SomeClass::someMeth(int, sig, *void)",
		"Subroutine", "SomeModule::SomeClass::someMeth(int,sig,*void)")

	// excessive space:
	parseAndAssert(t, "#0 SomeModule::SomeClass::someMeth(int,   sig=asd,    *void) (asd asd)",
		"Subroutine", "SomeModule::SomeClass::someMeth(int,sig=asd,*void)")
}

func TestCaptureArguments(t *testing.T) {
	parseAndAssert(t, "#0 ?? ()", "Arguments", "()")
	parseAndAssert(t, "#0  m4_traceon (obs=0x24eb0, argc=1, argv=0x2b8c8)", "Arguments", "(obs=0x24eb0,argc=1,argv=0x2b8c8)")
	parseAndAssert(t, "#1  0x68 in expand_macro (sym=0x2b) at mo.c:242", "Arguments", "(sym=0x2b)")
	parseAndAssert(t, "#2  0x40 in expand_token (obs=0x0, t=17, tdxb08)", "Arguments", "(obs=0x0,t=17,tdxb08)")

	// with excessive spaces
	parseAndAssert(t, "#0 ?? (   )", "Arguments", "()")
	parseAndAssert(t, "#0  m4_traceon ( obs= 0x24eb0,   argc =1,   argv=0x2b8c8 )", "Arguments", "(obs=0x24eb0,argc=1,argv=0x2b8c8)")
	parseAndAssert(t, "#1  0x68 in expand_macro (  sym= 0x2b  ) at mo.c:242", "Arguments", "(sym=0x2b)")
	parseAndAssert(t, "#2  0x40 in expand_token ( obs= 0x0, t = 17,  tdxb08)", "Arguments", "(obs=0x0,t=17,tdxb08)")

	// with excessive ()
	parseAndAssert(t, "#4  0x00007fdef42ddbcd in IlmThread_2_2::Semaphore::wait() () from /usr/lib/x86_64-linux-gnu/libIlmThread-2_2.so.12",
		"Arguments", "()")
}

func TestCaptureBinaryFile(t *testing.T) {
	// with excessive ()
	parseAndAssert(t, "#4  0x00007fdef42ddbcd in IlmThread_2_2::Semaphore::wait() () from /usr/lib/x86_64-linux-gnu/libIlmThread-2_2.so.12",
		"BinaryFile", "/usr/lib/x86_64-linux-gnu/libIlmThread-2_2.so.12")
}

func TestCaptureSourceFile(t *testing.T) {
	parseAndAssert(t, "#6  0x00007fdefb83c6ba in start_thread (arg=0x7fded47f8700) at pthread_create.c:333",
		"SourceFile", "pthread_create.c")
	parseAndAssert(t, "#0  0x00007fdefb844827 in futex_abstimed_wait_cancelable (private=0, abstime=0x0,  expected=0, futex_word=0x7fded989c3c8) at ../sysdeps/unix/sysv/linux/futex-internal.h:205",
		"SourceFile", "../sysdeps/unix/sysv/linux/futex-internal.h")
	parseAndAssert(t, "#2  0x00007fdefb8448d4 in __new_sem_wait_slow (sem=0x7fded989c3c8, abstime=0x0) at sem_waitcommon.c:181",
		"SourceFile", "sem_waitcommon.c")
}

func TestBlenderBackTrace(t *testing.T) {
	callback := func(line string, lineno int) error {
		if strings.HasPrefix(line, "Thread") {
			return nil
		}
		m, isEmptyLine, err := Parse(line)
		if err != nil {
			t.Errorf("Failed\ncan not parse line %d:\n%s\nreason: %s", lineno, line, err)
			return errors.New(line)
		}
		if !isEmptyLine {
			if len(m) == 0 {
				t.Log(line)
			}
			//t.Log(m)
		}
		return nil
	}
	err := iterLines("gdb.txt", callback)
	if err != nil {
		t.Fatal(err)
	}
}
