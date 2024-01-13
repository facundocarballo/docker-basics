import { Button, HStack, Input } from "@chakra-ui/react";
import React from "react";
import Task from "../types/task";

interface IInputTask {
  tasks: Task[],
  setTasks: (_tasks: Task[]) => void,
}

export const InputTask = ({tasks, setTasks}: IInputTask) => {
  // Attributes
  const [value, setValue] = React.useState<string>("");
  // Context
  // Methods
  const handleAddTask = async () => {
    const task = await Task.Create(value);
    if (task === undefined) {
      alert("Error creating this new task.");
      return
    }

    setTasks([...tasks, task]);
    setValue("");
  };
  // Component
  return (
    <>
      <HStack w={{ lg: "50%", sm: "90%" }}>
        <Input
          w="90%"
          placeholder="Insert a new task"
          value={value}
          onChange={(e) => setValue(e.currentTarget.value)}
          color='white'
        />
        <Button colorScheme="blue" onClick={handleAddTask} w='90px'>
          ADD
        </Button>
      </HStack>
    </>
  );
};
