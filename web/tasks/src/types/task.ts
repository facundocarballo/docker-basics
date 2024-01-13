import axios from "axios";
import { ConvertUint8ToDateString } from "../handlers/dates";
import { API, ApiTask } from "../handlers/api";

class Task {
  id: number;
  description: string;
  createdAt: string;

  constructor(_id: number, _description: string, _createdAt: Uint8Array) {
    this.id = _id;
    this.description = _description;
    this.createdAt = ConvertUint8ToDateString(_createdAt);
  }

  static async Create(_description: string): Promise<Task | undefined> {
    try {
      const { data } = await axios.post<ApiTask>(
        API.BaseUrl() + "task",
        JSON.stringify({
          description: _description,
        })
      );
      const task = new Task(data.id, data.description, data.createdAt);
      return task;
    } catch (err) {
      console.error("ERROR Task.Create function: ", err);
    }
  }

  static async GetAllTasks(): Promise<Task[]> {
    let tasks: Task[] = [];
    try {
      const { data } = await axios.get<ApiTask[]>(API.BaseUrl() + "task");
      data.forEach((task) =>
        tasks.push(new Task(task.id, task.description, task.createdAt))
      );
    } catch (err) {
      console.error("ERROR at Task.GetAllTasks function: ", err);
    }

    return tasks;
  }

  async Delete(): Promise<boolean> {
    try {
      await axios.delete(API.BaseUrl() + "task?id=" + this.id.toString());
    } catch (err) {
      console.error("ERROR at Task.Delete function: ", err);
      return false;
    }
    return true;
  }
}

export default Task;
