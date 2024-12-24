package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
 
	var n, m int
	fmt.Fscanf(in, "%d %d", &n, &m)

	result := fib(n, m)
	
	fmt.Fprintf(out, "%d\n", result)
}

func fib(n int, m int) int {
	reminders := make([]int, 2)
	reminders[0] = 0
	reminders[1] = 1
	period := 2
	periodDetected := false

	for i := 2; i <= n; i++ {
		next := (reminders[period-2] % m + reminders[period-1] % m) % m
		if next == 1 && reminders[period-1] == 0 {
			periodDetected = true
			period--
			break
		} else {
			reminders = append(reminders, next)
			period++
		}
	}

	if periodDetected {
		return reminders[n % period]
	} else {
		return reminders[period-1]
	}
}
