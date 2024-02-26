package luhn

import (
    "strings"
)

func Valid(id string) bool {
    idt := strings.TrimSpace(id)
    length := len(idt)
    if length <= 1 {
        return false
    }
    
    sum := 0
    count := 0
    for i := length - 1; i >= 0; i-- {
        digit := idt[i]
        if digit == ' ' {
            continue
        }
        if digit < '0' || digit > '9' {
            return false
        }
        d := int(digit - '0')
        if count%2 == 1 {
            d *= 2
            if d > 9 {
                d -= 9
            }
        }
        sum += d
        count++
    }
    
    return sum%10 == 0
}