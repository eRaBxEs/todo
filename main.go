package main

import (
	"fmt"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Status      int
}

type Project struct {
	ID    int
	Name  string
	Tasks []*Task //a container holding references to type Task
}

var ProjectList []*Project //a container holding references to type Project

func main() {
	// ProjectList = []*Project{} //

	p := newProject(&ProjectList, "proj 1")
	newTask(p, "task 1", "does something 1")
	newTask(p, "task 2", "does something 2")
	newTask(p, "task 3", "does something 3")

	p = newProject(&ProjectList, "proj 2")
	newTask(p, "task 1", "does something 1")
	newTask(p, "task 2", "does something 2")

	showProjects(&ProjectList)

	searchProjects(&ProjectList, "proj 3")
}

func showProjects(projs *[]*Project) { //a pointer to a container holding pointer to Projects
	for _, p := range *projs {
		fmt.Printf("\n[%d] Project %s | has %d tasks\n", p.ID, p.Name, len(p.Tasks))

		for _, i := range p.Tasks {
			fmt.Println(i.ID, i.Title, i.Description)
		}
	}
}

func newProject(plist *[]*Project, name string) (proj *Project) {
	proj = &Project{Name: name}
	proj.ID = len(*plist) + 1

	*plist = append(*plist, proj)
	return
}

func searchProjects(projlist *[]*Project, projname string) {
	counter := 0
	for _, p := range *projlist {
		if p.Name == projname {
			counter++
			fmt.Printf("The project name %q exist, its task are below:\n", projname)
			for _, i := range p.Tasks {
				fmt.Println(i.ID, i.Title, i.Description)
			}
			if counter == 1 {
				break
			} else {
				continue
			}

		}
	}
	if counter == 0 {
		fmt.Printf("The project name %q does not exist, so no task.Check project name\n", projname)
	}

}

//func searchTask()

func newTask(project *Project, title, descr string) (id int) {
	task := &Task{}
	task.ID = len(project.Tasks) + 1
	task.Title = title
	task.Description = descr

	project.Tasks = append(project.Tasks, task)

	return task.ID
}
