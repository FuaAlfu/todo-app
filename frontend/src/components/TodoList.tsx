import { Todo } from "../entities/Todo"
import { TodoItem } from "./TodoItem";


type Props = {
    todos: Todo[]
}
// FC = functional component
export const TodoList: React.FC<Props> = ({todos}) => {
return (
    <ul className="todo-list">
        {
            //i = key prop :: index
            todos.map((todo, i) => (
                <li key={i}><TodoItem 
                title={todo.title}
                description={todo.description}
                 isCompleted={todo.isCompleted}
                 /></li>
))
        }
    </ul>
    )
}