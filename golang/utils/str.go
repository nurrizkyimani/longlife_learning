
package utils

import ( 
  "github.com/nurrizkyimani/longlife_learning/golang/greetings"
)


func TestingLen(inp_str string) int {
  mn := len(inp_str)
  return mn
}  


func HellosV3(names []string,) (map[string]string, error) {

  messages := make(map[string]string)

  for _, name := range names {
    message, err := greetings.HelloV2(name)
    if err != nil {
      return messages, err
    }
    messages[name] = message
  }

  return messages, nil
}