import React from 'react';
import './styles/app.scss';
import {Header} from './components/Header';
import {TodoList} from './components/TodoList';

function App() {
  return (
    <div className="App">
     <Header />
     <TodoList todos={[
       {title: "keep programming", isCompleted: false},
       {title: "learning",description: 'learning about new things', isCompleted: false}
       ]} />
    </div>
  );
}

export default App;
