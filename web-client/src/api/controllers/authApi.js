import http from "../http";
import jwt from "../../utils/jwt";

class AuthApi {
    async getUser() {
        try {
            const {data} = await http.get('/user');
            console.log('get user', data)
            return data;
        } catch (err) {
            console.log(err)
        }
    }

    async login(values) {
        try {
            console.log('login')
            const {data} = await http.post('/login', values);
            jwt.set(data.jwt);
            return this.getUser();
        } catch (err) {
            console.log(err)
        }
    }

    async signIn(values) {
        try {
            console.log('signIn')
            const {data} = await http.post('/signIn', values);
            jwt.set(data.jwt);
            return this.getUser();
        } catch (err) {
            console.log(err)
        }
    }

    async signOut() {
        try {
            const res = await http.post('/signOut');
            jwt.remove();
            console.log('signOut', res)
            return res;
        } catch (err) {
            console.log(err)
        }
    }
}

export default new AuthApi();