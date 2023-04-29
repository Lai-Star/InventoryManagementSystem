import axios from 'axios';

const loginRoute = async (username, password) => {
    return await axios.post("http://localhost:8080/login", {
        username, password
    });
}

export default loginRoute;