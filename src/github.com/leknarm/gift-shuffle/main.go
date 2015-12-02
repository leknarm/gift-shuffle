package main
 
import (
   "math/rand"
   "time"
// "net/smtp"
   "encoding/json"
   "io/ioutil"
   "net/mail"
   "strings"
   "fmt"
   "log"
   "os"
)

type User struct {
   Email string
   Name  string
}

type Users struct {
   Users []User
}
 
func encodeRFC2047(String string) string{
   // use mail's rfc2047 to encode any string
   addr := mail.Address{String, ""}
   return strings.Trim(addr.String(), " <>")
}
 
func main() {
   var arg = os.Args[1]
   fmt.Println(arg)

   var u Users
   var m = make(map[string]string)

   file, err := ioutil.ReadFile("/Users/Tantai/GitHub/src/github.com/leknarm/gift-shuffle/template.json")
   if err != nil {
      log.Fatal(err)
   }
   
   err = json.Unmarshal(file, &u)
   if err != nil {
      log.Fatal(err)
   }

   for _, user := range u.Users {
      m[user.Email] = user.Name
   }

   var keys = make([]string, 0, len(m))
   for k := range m {
      keys = append(keys, k)
   }

   var members = shuffle(keys)

   // auth := smtp.PlainAuth("", "xxx@leknarm.com", "xxx", "smtp.gmail.com")
   // subject := "Subject: Test New Year Cerebration\r\n\r\n"
 
   for i, _ := range members {
      if i != len(members) - 1 {
         message := "Hi,\r\n\r\nYour buddy is " + m[members[i]] + "\r\n\r\nPlease by a gift price > 550 baht and send to your buddy on new year party. So please keep this secret.!!!"
         fmt.Println("\n\nsend mail to " + members[i+1] + " with message: " + message)
         // smtp.SendMail("smtp.gmail.com:587", auth, "xxx@leknarm.com", []string{members[i+1]}, []byte(subject + message))
      }
   }
   message := "Hi,\r\n\r\nYour buddy is " + m[members[len(members) - 1]] + "\r\n\r\nPlease by a gift price > 550 baht and send to your buddy in the new year party. So please keep this secret.!!!"
   fmt.Println("\n\nsend mail to " + members[0] + "with message: " + message)
   // smtp.SendMail("smtp.gmail.com:587", auth, "xxx@leknarm.com", []string{members[0]}, []byte(subject + message))
}
 
func shuffle(members []string) []string {
   rand.Seed(time.Now().UTC().Unix())
   var newMembers = make([]string, len(members))
   var rands = rand.Perm(len(members))
   for i, v := range rands {
      newMembers[i] = members[v]
   }
   return newMembers
}