import { Component, OnInit } from '@angular/core';
import { DebitService } from '../../services/debit.service';

import { Debit } from '../../models/Debit';

@Component({
  selector: 'app-debits',
  templateUrl: './debits.component.html',
  styleUrls: ['./debits.component.css'],
})
export class DebitsComponent implements OnInit {
  debits: Debit[];

  constructor(private debitService: DebitService) {}

  ngOnInit(): void {
    this.debitService.getDebits().subscribe((debits) => {
      debits.forEach((current, index, array) => {
        array[index].postedDate = new Date(current.postedDate);
      });
      this.debits = debits;
    });
  }

  deleteDebit(debit: Debit) {
    // Remove from UI
    this.debits = this.debits.filter((t) => t.id !== debit.id);
    // Remove from server
    this.debitService.deleteDebit(debit).subscribe();
  }

  addDebit(debit: Debit) {
    this.debitService.addDebit(debit).subscribe((newDebit) => {
      newDebit[0].postedDate = new Date(newDebit[0].postedDate);
      this.debits.push(newDebit[0]);
    });
  }
}
