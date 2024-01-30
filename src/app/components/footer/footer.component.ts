import { Component } from '@angular/core';

@Component({
  selector: 'app-footer',
  standalone: true,
  imports: [],
  templateUrl: './footer.component.html',
  styleUrl: './footer.component.css',
})
export class FooterComponent {
  public logo: string = '/assets/images/logo.svg';
  public title: string = 'Footer';
  public styles = {
    container: 'flex flex-row items-center bg-sky-50 p-2',
  };
}
