import { Component } from '@angular/core';

@Component({
  selector: 'app-header',
  standalone: true,
  imports: [],
  templateUrl: './header.component.html',
  styleUrl: './header.component.css',
})
export class HeaderComponent {
  public logo: string = '/assets/images/logo.svg';
  public title: string = 'My library';
  public styles = {
    container: 'flex flex-row items-center bg-sky-50 p-2',
  };
}
