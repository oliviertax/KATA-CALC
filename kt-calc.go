package main  //go 1.21.3
 
import ( 
 "bufio" 
 "fmt" 
 "os" 
 "strconv" 
 "strings" 
) 
 
func main() { 
 
 reader, _ := bufio.NewReader(os.Stdin).ReadString('\n') 
 var x int 
 var result int 
 reader = strings.ReplaceAll(reader, " ", "") 
 for i:=0;i<=len(reader);i++ { 
  if string(reader[i])=="+" || string(reader[i])=="-" || string(reader[i])=="*" || string(reader[i])=="/"  { 
    x=i 
    break 
   } else if strings.Count(string(reader), "+")==2 || strings.Count(string(reader), "-")==2 || strings.Count(string(reader), "*")==2 || strings.Count(string(reader), "/")==2 {
     panic("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
  } else if i==len(reader)-1{  
     panic("Вывод ошибки, так как строка не является математической операцией.")  
  } 
  
 } 
  
 a:=romanToArabic(reader[0:x]) 
 b:=romanToArabic(reader[x+1:len(reader)-2]) 
 if string(reader[x])=="-" && a<b {
    panic("Римское число не может быть отрицательным")
	}
  
 for i:=0;i<=9;i++ { 
  if reader[0:x]==strconv.Itoa(i) || reader[0:x]=="10" { 
   result=calc(reader) 
  }  
 } 
    for i:=48;i<=57;i++{ 
        if reader[0]==byte(i){ 
            fmt.Println(result) 
            break 
        } else if i==57 && result==0{ 
          result=romancalc(a,b,string(reader[x]),reader) 
            fmt.Println(Roman(result)) 
        } 
    } 
 } 
 
func plus(a, b int) int { 
 return a + b 
} 
 
func minus(a, b int) int { 
 return a - b 
} 
 
func umnozhenie(a, b int) int { 
 return a * b 
} 
 
func delenie(a, b int) int { 
 return a / b 
} 
 
 
func Roman(number int) string{ 
    rimskie := []struct{ 
        val int 
        chislo string 
    }{ 
        {100, "C"}, 
        {90, "XC"}, 
        {50, "L"}, 
        {40, "XL"}, 
        {10, "X"}, 
        {9, "IX"}, 
        {8, "VIII"}, 
        {7, "VII"}, 
        {6, "VI"}, 
        {5, "V"}, 
        {4, "IV"}, 
        {3, "III"}, 
        {2, "II"}, 
        {1, "I"}, 
    } 

    roman := "" 
    for _, conv := range rimskie { 
        for number >= conv.val { 
            roman += conv.chislo 
            number -= conv.val 
        } 
    } 
 return roman 
} 
 
func romanToArabic(roman string) int { 

 switch roman { 
 case "I":  
    return 1 
 case "II": 
    return 2 
 case "III": 
    return 3 
 case "IV": 
    return 4 
 case "V": 
    return 5 
 case "VI": 
    return 6 
 case "VII": 
    return 7 
 case "VIII": 
    return 8 
 case "IX": 
    return 9 
 case "X": 
    return 10 
 } 
 return 0
} 
 
func romancalc(a,b int,operator,reader string) int { 
    reader = strings.ReplaceAll(reader, " ", "") 

 switch { 
 case operator == "+" : 
    return plus(a, b) 
 
 case operator == "-" : 
    return minus(a, b) 
 
 case operator == "*" : 
    return umnozhenie(a, b) 
 
 case operator == "/" : 
     return delenie(a, b) 
 } 
    return 1 
} 
 
func calc(reader string) int { 
 var str string 
    reader = strings.ReplaceAll(reader, " ", "") 
    a, _ := strconv.Atoi(string(reader[0])) 
     b, _ := strconv.Atoi(string(reader[2])) 
    str=string(reader[1]) 
  
    if reader[1]>48{
        panic("Число не должно превышать десяти")
     }
  
 if reader[1] == 48 { 
    a = 10 
    str=string(reader[2]) 
    b, _ = strconv.Atoi(string(reader[3])) 

  if reader[4] == 48 { 
     b = 10 
  } 
 
 } else if reader[3] == 48 { 
     b = 10 
     str=string(reader[1]) 
    a, _ = strconv.Atoi(string(reader[0])) 
 }  
 
 switch { 
 case str == "+" : 
     return plus(a, b) 
 
 case str == "-" : 
     return minus(a, b) 
 
 case str == "*" : 
    return umnozhenie(a, b) 
 
 case str == "/" : 
    return delenie(a, b) 
 } 
    return 0
}