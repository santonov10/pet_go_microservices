import axios from "axios";

class Http {
    constructor() {
        this.axios = axios.create({
            baseURL: '/api/v1',
            headers: {
                'Authorization': `Bearer ${localStorage.getItem('jwt')}`
            }
        })
    }

    get(url, config = {}) {
        return this.axios.get(url, config)
    }

    post(url, body, config = {}) {
        return this.axios.post(url, body, config)
    }
}

export default new Http();