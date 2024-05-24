import http from 'k6/http';

export const GetAllTasks = () => {
    http.get('http://localhost:8000/tasks')


}
