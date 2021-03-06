import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Debit } from '../models/Debit';
import { environment } from '../../environments/environment';

const httpHeaders = new HttpHeaders({
  'Content-Type': 'application/json', // eslint-disable-line @typescript-eslint/naming-convention
});

@Injectable({
  providedIn: 'root',
})
export class DebitService {
  serverUrl: string = environment.serverURL + '/debits';

  constructor(private http: HttpClient) {}

  // Get all Debits
  getDebits(): Observable<Debit[]> {
    const url = `${this.serverUrl}`;
    return this.http.get<Debit[]>(url, {headers: httpHeaders});
  }

  // Delete Debit
  deleteDebit(debit: Debit): Observable<Debit> {
    const url = `${this.serverUrl}/${debit.id}`;
    return this.http.delete<Debit>(url);
  }

  // Add Debit
  addDebit(debit: Debit): Observable<Debit> {
    return this.http.post<Debit>(this.serverUrl, debit, {
      headers: httpHeaders,
    });
  }

  // Update Debit
  toggleCompleted(debit: Debit): Observable<any> {
    return this.http.put(this.serverUrl, debit, { headers: httpHeaders });
  }
}
