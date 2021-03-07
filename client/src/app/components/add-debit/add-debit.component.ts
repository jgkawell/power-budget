import { Component, OnInit, EventEmitter, Output } from '@angular/core';

import { Debit } from '../../models/Debit';

@Component({
  selector: 'app-add-debit',
  templateUrl: './add-debit.component.html',
  styleUrls: ['./add-debit.component.css'],
})
export class AddDebitComponent implements OnInit {
  @Output() addDebit: EventEmitter<any> = new EventEmitter();

  vendor: string;

  constructor() {}

  ngOnInit(): void {}

  onSubmit() {
    const debit: Debit = {
      posted_date: new Date(),
      amount: 1.99,
      vendor: this.vendor,
      purpose: 'Groceries',
      account_id: '422f729f-8833-4be0-a6d8-ba980710c2de',
      budget: 1,
      notes: 'This is a note',
    };

    this.addDebit.emit(debit);
    this.vendor = null;
  }
}
