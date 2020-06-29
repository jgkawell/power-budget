import { Component, OnInit, EventEmitter, Output } from '@angular/core';

import { Todo } from '../../models/Todo';

@Component({
  selector: 'app-add-todo',
  templateUrl: './add-todo.component.html',
  styleUrls: ['./add-todo.component.css'],
})
export class AddTodoComponent implements OnInit {
  @Output() addTodo: EventEmitter<any> = new EventEmitter();

  title: string;

  constructor() {}

  ngOnInit(): void {}

  onSubmit() {
    const todo: Todo = {
      title: this.title,
      completed: false,
    };

    this.addTodo.emit(todo);
    this.title = null;
  }
}
