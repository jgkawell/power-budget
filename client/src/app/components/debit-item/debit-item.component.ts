import { Component, OnInit, Input, EventEmitter, Output } from '@angular/core';
import { DebitService } from '../../services/debit.service';

import { Debit } from 'src/app/models/Debit';

@Component({
  selector: 'app-debit-item',
  templateUrl: './debit-item.component.html',
  styleUrls: ['./debit-item.component.css'],
})
export class DebitItemComponent implements OnInit {
  @Input() debit: Debit;
  @Output() deleteDebit: EventEmitter<Debit> = new EventEmitter();

  isComplete = false;

  constructor(private debitService: DebitService) {}

  ngOnInit(): void {
    // Required by angular
  }

  // Set Dynamic Classes
  setClasses() {
    return {
      debit: true,
      'is-complete': this.isComplete,
    };
  }

  onToggle(debit: Debit) {
    this.isComplete = !this.isComplete;
  }

  onDelete(debit: Debit) {
    this.deleteDebit.emit(debit);
  }
}
