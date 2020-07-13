import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Debit } from '../models/Debit';
import { environment } from '../../environments/environment';

const httpHeaders = new HttpHeaders({
  'Content-Type': 'application/json',
});

@Injectable({
  providedIn: 'root',
})
export class DebitService {
  serverUrl: string = environment.serverURL + '/debit';

  constructor(private http: HttpClient) {}

  // Get all Debits
  getDebits(): Observable<Debit[]> {
    const url = `${this.serverUrl}/all`;
    return this.http.get<Debit[]>(url);
  }

  // Delete Debit
  deleteDebit(debit: Debit): Observable<Debit> {
    const url = `${this.serverUrl}`;
    const options = {
      headers: httpHeaders,
      body: { id: debit.id },
    };
    return this.http.delete<Debit>(url, options);
  }

  // Add Debit
  addDebit(debit: Debit): Observable<Debit> {
    return this.http.post<Debit>(this.serverUrl, debit, { headers: httpHeaders});
  }

  // Update Debit
  toggleCompleted(debit: Debit): Observable<any> {
    return this.http.put(this.serverUrl, debit, { headers: httpHeaders});
  }
}
