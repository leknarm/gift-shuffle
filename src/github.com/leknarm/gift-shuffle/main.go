package main
 
import (
   "math/rand"
   "time"
   "net/smtp"
   "encoding/json"
   "io/ioutil"
   "net/mail"
   "strings"
   "fmt"
   "log"
   "os"
   "github.com/howeyc/gopass"
)

type User struct {
   Email string
   Name  string
}

type Users struct {
   Subject  string
   Message  string
   Users    []User
}
 
func encodeRFC2047(String string) string{
   // use mail's rfc2047 to encode any string
   addr := mail.Address{String, ""}
   return strings.Trim(addr.String(), " <>")
}
 
func main() {
   var u Users
   var m = make(map[string]string)
   var email string
   var host = "smtp.gmail.com:587"

   fmt.Println(len(os.Args))

   if len(os.Args) == 1 {
      log.Fatalf("Not found template file argument.")
   }

   fmt.Print("Email: ")
   _, err := fmt.Scanf("%s\n", &email)
   if err != nil {
      log.Fatal(err)
   }

   fmt.Print("Password: ")
   pass := gopass.GetPasswd()

   arg := os.Args[1]
   

   file, err := ioutil.ReadFile(arg)
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

   auth := smtp.PlainAuth("", email, string(pass), "smtp.gmail.com")
   subject := u.Subject
 
   for i, _ := range members {
      if i != len(members) - 1 {
         message := fmt.Sprintf(u.Message, m[members[i+1]], m[members[i]])
         fmt.Println(message)
         smtp.SendMail(host, auth, email, []string{members[i+1]}, []byte(subject + message))
      }
   }
   message := fmt.Sprintf(u.Message, m[members[0]], m[members[len(members) - 1]])
   fmt.Println(message)
   smtp.SendMail(host, auth, email, []string{members[0]}, []byte(subject + message))
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