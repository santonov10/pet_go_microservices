import {Link} from "react-router-dom";

export default function MainPage() {
    return (
        <div className="MainPage">
            <h3>Сервис для создания задач</h3>
            <Link to="/tasks">Перейти к задачам</Link>
        </div>
    );
}