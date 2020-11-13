package main 
  
import "fmt"
import "strings"
import "strconv"
import "log"



// 12/14/1989

type DOB struct {
  Month string
  Day   string
  Year  string
}

var dob     string
var intData []int
var warn    string = "ERROR: Input is not valid.\n\nYou entered:"

func clear() {
  fmt.Print("\033[H\033[2J")
}

func warnUser(warning string, blank int) {
  if blank == 0 {
    clear()
    fmt.Printf("%s", warning)
  } else {
    clear()
    fmt.Printf("ERROR: Input is not valid.\n%s\n", warning) 
  }

}

func convStrsToInts(a []string) []int {
  z := []int{}
  for k, _ := range a {
    newint, err := strconv.Atoi(a[k])
    if err != nil {
      log.Fatal(err)
    }

    z =  append(z, newint)
  }
  
  return z
}

func updateDOB(updateDOB []string) {
  dob = fmt.Sprintf("%s/%s/%s", updateDOB[0], updateDOB[1], updateDOB[2])
}

func validateDobInput(s string) bool {
  if len(s) < 1 {
    warnUser("", 0)
    return false
  }

  data := strings.Split(s, "/")

  if len(data) != 3 {
    warnUser("\nToo short. Please provide full DOB.\n", 1)
    return false
  }

  if len(data[0]) != 2 {
    //convert variable 'data' to []int and check ints for validity
    intData = convStrsToInts(data)

    //if the number in question is between 1&9, add the zero for use with time.Time obj
    if intData[0] < 1 || intData[0] > 9 {
      warnUser(fmt.Sprintf("\nMonth provided was not correct. You entered: %s\n", s), 1)
      return false
    } else {
      data[0] = "0" + data[0]
    }
  }

  if len(data[1]) != 2 {
    intData = convStrsToInts(data)
    if intData[1] < 1 || intData[1] > 9 {
      warnUser(fmt.Sprintf("\nDay should be 2 digits. You entered: %s\n", s), 1)
      return false
    } else {
      data[1] = "0" + data[1]
    }
  }

  if len(data[2]) != 4 {
    clear()
    fmt.Println(warn, "\nYear should be 4 digits\n")
    return false
  }

  updateDOB(data)
  return true
}

func gatherDOB() {
  fmt.Println("With or without slashes, please enter your birthday (mm/dd/yyyy)")
  fmt.Scanln(&dob) 
}

func main() {
  //Initial attempt to gather DOB
  clear()
  gatherDOB()

  //If initial attempt fails, loop until we get it right
  for validateDobInput(dob) != true {
    gatherDOB()
  }
  fmt.Printf("You entered: %s\n", dob)
}
