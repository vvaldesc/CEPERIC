import { ApplicationConfig, provideZoneChangeDetection } from '@angular/core';
import { provideRouter } from '@angular/router';

import { routes } from './app.routes';
import { provideAnimationsAsync } from '@angular/platform-browser/animations/async';
import { initializeApp, provideFirebaseApp } from '@angular/fire/app';
import { getAuth, provideAuth } from '@angular/fire/auth';
import { getFirestore, provideFirestore } from '@angular/fire/firestore';
import { getDatabase, provideDatabase } from '@angular/fire/database';
import { getPerformance, providePerformance } from '@angular/fire/performance';
import { getStorage, provideStorage } from '@angular/fire/storage';

export const appConfig: ApplicationConfig = {
  providers: [
    provideZoneChangeDetection({ eventCoalescing: true }), 
    provideRouter(routes), 
    provideAnimationsAsync(), 
    provideFirebaseApp(() => initializeApp({
      projectId: "cepe-test",
      appId: "1:803389290966:web:56f59a93a6c7438e128c77",
      storageBucket: "cepe-test.firebasestorage.app",
      apiKey: "AIzaSyBHBvvS57_JM1-Cc0ZqMtCCN6LcA8YiS5c",
      authDomain: "cepe-test.firebaseapp.com",
      messagingSenderId: "803389290966",
      measurementId: "G-Y51V81PD87"
    })), 
    provideAuth(() => getAuth()), 
    provideFirestore(() => getFirestore()), 
    provideDatabase(() => getDatabase()), 
    providePerformance(() => getPerformance()), 
    provideStorage(() => getStorage())
  ]
};
