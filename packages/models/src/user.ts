export interface User {
  id: number;
  name: string;
  email: string;
}

export interface UserData extends User {
  _id: string;
}
