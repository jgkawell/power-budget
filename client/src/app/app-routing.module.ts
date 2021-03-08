import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { DebitsComponent } from './components/debits/debits.component';
import { AboutComponent } from './components/pages/about/about.component';

const routes: Routes = [
  {
    path: '',
    component: DebitsComponent,
  },
  {
    path: 'about',
    component: AboutComponent,
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
