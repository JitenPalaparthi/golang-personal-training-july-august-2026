package main

func main() {
	num := 1
loop:
	if num <= 10 {
		if num%2 == 0 {
			println("even number:", num)
		}
		num += 1
		goto loop
	}

	num = 1

until:
	if num <= 10 {
		if num%2 == 0 {
			goto even
		} else {
			goto odd
		}
	} else {
		goto exit
	}

even:
	println(num, "is even number")
	num++
	goto until

odd:
	println(num, "is odd number")
	num++
	goto until
exit:
}
