package main 
import pqueue "github.com/nu7hatch/gopqueue" 

type Task struct { 
  Name string 
  priority int 
}

func (t *Task) Less(other interface{}) bool { 
  return t.priority > other.(*Task).priority 
}

func main() { 
  q := pqueue.New(0) 
  q.Enqueue(&Task{"Victoria", 46}) 
  q.Enqueue(&Task{"Honduras Progreso", 38}) 
  q.Enqueue(&Task{"Lobos UPNFM", 42}) 
  q.Enqueue(&Task{"Marathon", 50}) 
  q.Enqueue(&Task{"Motagua", 63}) 
  q.Enqueue(&Task{"Olimpia", 67}) 
  q.Enqueue(&Task{"Platense", 29}) 
  q.Enqueue(&Task{"Real España", 75}) 
  q.Enqueue(&Task{"Real Sociedad", 33}) 
  q.Enqueue(&Task{"Vida", 61})

  println("Tabla de Posiciones")
  for i := 0; i < 10; i += 1 {
    task := q.Dequeue()
    println(task.(*Task).Name)
    if i == 9 {
      println()
		  print("El equipo que desciende a segunda division es: ")
		  println(task.(*Task).Name)
	  }
  }
}
