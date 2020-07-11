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
      postedDate: new Date(),
      amount: 0,
      vendor: this.vendor,
      purpose: 'temp',
      account: 'temp',
      budget: 0,
      notes: 'temp',
    };

    this.addDebit.emit(debit);
    this.vendor = null;
  }
}
