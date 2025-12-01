// Ejemplo de servicio Firebase - src/app/services/firebase.service.ts
// Para usar este servicio, descomenta y ajusta según tus necesidades

/*
import { Injectable } from '@angular/core';
import { Firestore, collection, collectionData, doc, setDoc, deleteDoc } from '@angular/fire/firestore';
import { Observable } from 'rxjs';

export interface Item {
  id?: string;
  name: string;
  description: string;
  createdAt: Date;
}

@Injectable({
  providedIn: 'root'
})
export class FirebaseService {
  constructor(private firestore: Firestore) {}

  // Obtener todos los items
  getItems(): Observable<Item[]> {
    const itemsCollection = collection(this.firestore, 'items');
    return collectionData(itemsCollection, { idField: 'id' }) as Observable<Item[]>;
  }

  // Agregar un item
  async addItem(item: Item): Promise<void> {
    const itemDoc = doc(this.firestore, `items/${Date.now()}`);
    await setDoc(itemDoc, {
      ...item,
      createdAt: new Date()
    });
  }

  // Eliminar un item
  async deleteItem(id: string): Promise<void> {
    const itemDoc = doc(this.firestore, `items/${id}`);
    await deleteDoc(itemDoc);
  }
}
*/

// Ejemplo de uso en un componente:
/*
import { Component, OnInit } from '@angular/core';
import { FirebaseService, Item } from './services/firebase.service';
import { CommonModule } from '@angular/common';
import { MatCardModule } from '@angular/material/card';
import { MatButtonModule } from '@angular/material/button';

@Component({
  selector: 'app-items',
  standalone: true,
  imports: [CommonModule, MatCardModule, MatButtonModule],
  template: `
    <div class="p-4">
      <h2 class="text-2xl font-bold mb-4">Items de Firebase</h2>
      
      @for (item of items$ | async; track item.id) {
        <mat-card class="mb-4">
          <mat-card-header>
            <mat-card-title>{{ item.name }}</mat-card-title>
          </mat-card-header>
          <mat-card-content>
            <p>{{ item.description }}</p>
          </mat-card-content>
          <mat-card-actions>
            <button mat-button color="warn" (click)="deleteItem(item.id!)">
              Eliminar
            </button>
          </mat-card-actions>
        </mat-card>
      }
    </div>
  `
})
export class ItemsComponent implements OnInit {
  items$!: Observable<Item[]>;

  constructor(private firebaseService: FirebaseService) {}

  ngOnInit() {
    this.items$ = this.firebaseService.getItems();
  }

  async addItem() {
    await this.firebaseService.addItem({
      name: 'Nuevo Item',
      description: 'Descripción del item',
      createdAt: new Date()
    });
  }

  async deleteItem(id: string) {
    await this.firebaseService.deleteItem(id);
  }
}
*/
