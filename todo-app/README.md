
# To-Do CLI

A simple and efficient command-line to-do list application built with Go. Manage your tasks directly from the terminal without leaving your workflow.

---

## Features

- **Add** new tasks to your list.
- **List** all pending tasks in a clean, readable format.
- **Update** existing tasks.
- **Delete** completed or unnecessary tasks.

---

## Usage

The application is controlled via command-line flags.


Usage of todo.exe:
-a string
add todo
-d string
delete todo
-l    List todo
-u string
update todo


### Commands

-   **Add a new to-do:**
    -   Flag: `-a "<new todo to be added>"`
    -   The task description must be enclosed in double quotes.

-   **List all to-dos:**
    -   Flag: `-l`
    -   Displays all tasks with their corresponding numbers.

-   **Update an existing to-do:**
    -   Flag: `-u "<task number>,<updated todo>"`
    -   Specify the task number to update, followed by a comma, and the new task description. The entire argument must be enclosed in double quotes.

-   **Delete a to-do:**
    -   Flag: `-d <task number>`
    -   Specify the number of the task you want to remove.

---

## Examples

### 1. Add a new task

```sh
$ ./todo.exe -a "buy milk"
```sh
$ ./todo.exe -a "buy fruits"

2. List all tasks
$ ./todo.exe -l

Output:

No. | Todo
----|-------------------
1   | buy milk
2   | buy fruits

3. Update a task
To change task #1 from "buy milk" to "buy almond milk":

$ ./todo.exe -u "1,buy almond milk"

New List:

No. | Todo
----|-------------------
1   | buy almond milk
2   | buy fruits

4. Delete a task
To delete task #2:

$ ./todo.exe -d 2

New List:

No. | Todo
----|-------------------
1   | buy almond milk
