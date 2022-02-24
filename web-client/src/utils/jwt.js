const JWT_KEY = 'jwt';

class JWT {
    get() {
        return localStorage.getItem(JWT_KEY)
    }

    set(value) {
        localStorage.setItem(JWT_KEY, value)
    }

    remove() {
        localStorage.removeItem(JWT_KEY)
    }
}

export default new JWT();