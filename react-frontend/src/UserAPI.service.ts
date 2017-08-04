
import { User } from "./API_models";

const baseUri = 'http://localhost:8080';

class UserAPIService {
    GetAllUsers(): Promise<User[]> {
        return fetch(`${baseUri}/users`).then(resp => resp.json())
    }
}

export default new UserAPIService();