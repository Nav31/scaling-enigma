package codewars_probs

import (
    "strings"
)

// https://www.codewars.com/kata/convert-string-to-camel-case/train/go

func ToCamelCase(s string) string {
    var delimiter string
    retStr := ""
    if strings.Index(s, "-") > -1 {
        delimiter = "-"
    } else {
        delimiter = "_"
    }
    splitString := strings.Split(s, delimiter)
    for idx, val := range splitString {
       strVal := string(val)
       if idx > 0 {
           retStr += strings.ToUpper(strVal[0:1]) + strVal[1:]
       } else {
           retStr += strVal
       }
    }
    return retStr
}
