import Task from "../types/task";

export const DeleteTaskFromArray = (task: Task, tasks: Task[]): Task[] => {
    let nTasks: Task[] = []
    tasks.forEach(t => {
        if (t.id !== task.id) nTasks.push(t);
    })
    return nTasks;
}