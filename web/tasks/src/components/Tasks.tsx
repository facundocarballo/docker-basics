import React from "react";
import Task from "../types/task";
import { TaskComponent } from "./TaskComponent";

interface ITasks {
    tasks: Task[],
    setTasks: (_tasks: Task[]) => void,
}
export const Tasks = ({tasks, setTasks}: ITasks) => {
  // Attributes
  // Context
  // Methods
  // Component
  return (
    <>
      {tasks.map((task, idx) => (
        <TaskComponent
          task={task}
          setTasks={setTasks}
          tasks={tasks}
          key={idx}
        />
      ))}
    </>
  );
};
