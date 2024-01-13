import { InputTask } from "@/src/components/InputTask";
import { Tasks } from "@/src/components/Tasks";
import { Title } from "@/src/components/Title";
import Task from "@/src/types/task";
import { Divider, VStack, Box } from "@chakra-ui/react";
import Head from "next/head";
import React from "react";

export default function Home() {
  // Attributes
  const [tasks, setTasks] = React.useState<Task[]>([]);
  // Context
  // Methods
  const handleGetAllTasks = async () => {
    const tasks = await Task.GetAllTasks();
    setTasks(tasks);
  };
  // Component
  React.useEffect(() => {
    handleGetAllTasks();
  }, []);
  return (
    <>
      <Head>
        <title>Tasks</title>
        <meta name="description" content="Task handler to start using docker" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <VStack w="full">
        <Title />
        <InputTask setTasks={setTasks} tasks={tasks}/>
        <Box h="15px" />
        <Divider />
        <Box h="15px" />
        <Tasks setTasks={setTasks} tasks={tasks}/>
      </VStack>
    </>
  );
}
