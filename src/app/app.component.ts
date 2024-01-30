import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { FooterComponent } from './components/footer/footer.component';
import { HeaderComponent } from './components/header/header.component';

const Modules = [CommonModule, RouterOutlet];
const Components = [HeaderComponent, FooterComponent];

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [...Modules, ...Components],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css',
})
export class AppComponent {
  title = 'library';
}
