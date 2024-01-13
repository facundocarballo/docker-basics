import { Button, HStack, Text, Spacer, Box } from '@chakra-ui/react';
import React from 'react';
import Task from '../types/task';
import { DeleteTaskFromArray } from '../handlers/tasks';

interface ITaskComponent {
   task: Task,
   tasks: Task[],
   setTasks: (_tasks: Task[]) => void,
}
export const TaskComponent = ({task, tasks, setTasks}: ITaskComponent) => {
   // Attributes
   // Context
   // Methods
   const handleDelete = async () => {
      const nTasks = DeleteTaskFromArray(task, tasks);
      setTasks(nTasks);
      
      await task.Delete();
   }

   // Component
   return(
      <>
        <HStack w={{lg: '50%', sm: "95%", base:"95%"}} bg='blue.700' borderRadius={10}>
            <Box w='5px' />
            <Text fontSize='14px' color='white'>
               {task.description}
            </Text>
            <Box w='5px' />
            <Spacer />
            <Button colorScheme='red' onClick={handleDelete} w='90px'>
               DELETE
            </Button>
        </HStack>
      </>
   );
}