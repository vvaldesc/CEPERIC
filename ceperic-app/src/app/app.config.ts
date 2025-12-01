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
      apiKey: "AIzaSyARW0jAxCSWJgpFwCLJ5Sy_Zif02-FOo8E",
      authDomain: "ceperic-68bcd.firebaseapp.com",
      projectId: "ceperic-68bcd",
      storageBucket: "ceperic-68bcd.firebasestorage.app",
      messagingSenderId: "382073954600",
      appId: "1:382073954600:web:a2d919a22172001637301d",
      measurementId: "G-0JDN0G4K0L"
    })), 
    provideAuth(() => getAuth()), 
    provideFirestore(() => getFirestore()), 
    provideDatabase(() => getDatabase()), 
    providePerformance(() => getPerformance()), 
    provideStorage(() => getStorage())
  ]
};
