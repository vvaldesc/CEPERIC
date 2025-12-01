# Configuración de Firebase

## 📝 Pasos para configurar Firebase

### 1. Crear proyecto en Firebase Console
1. Ve a [Firebase Console](https://console.firebase.google.com/)
2. Crea un nuevo proyecto o selecciona uno existente
3. Registra tu aplicación web

### 2. Obtener credenciales
En la configuración del proyecto, encontrarás algo similar a:

```typescript
const firebaseConfig = {
  apiKey: "AIza...",
  authDomain: "tu-proyecto.firebaseapp.com",
  projectId: "tu-proyecto",
  storageBucket: "tu-proyecto.appspot.com",
  messagingSenderId: "123456789",
  appId: "1:123456789:web:abc123"
};
```

### 3. Configurar en la aplicación
Las credenciales ya están configuradas en `src/app/app.config.ts` por el comando `ng add @angular/fire`.

Si necesitas actualizarlas manualmente, edita ese archivo.

### 4. Habilitar servicios en Firebase Console

#### Firestore Database
1. Ve a "Firestore Database" en el menú lateral
2. Clic en "Crear base de datos"
3. Selecciona el modo (producción o prueba)
4. Elige la ubicación de los datos

#### Authentication (Opcional)
1. Ve a "Authentication" en el menú lateral
2. Clic en "Comenzar"
3. Habilita los proveedores que necesites (Email/Password, Google, etc.)

#### Storage (Opcional)
1. Ve a "Storage" en el menú lateral
2. Clic en "Comenzar"
3. Configura las reglas de seguridad

### 5. Configurar reglas de seguridad de Firestore

Para desarrollo, puedes usar reglas permisivas (NO para producción):

```javascript
rules_version = '2';
service cloud.firestore {
  match /databases/{database}/documents {
    match /{document=**} {
      allow read, write: if request.time < timestamp.date(2025, 12, 31);
    }
  }
}
```

Para producción, usa reglas más restrictivas:

```javascript
rules_version = '2';
service cloud.firestore {
  match /databases/{database}/documents {
    match /items/{itemId} {
      allow read: if true;
      allow write: if request.auth != null;
    }
  }
}
```

### 6. Uso en la aplicación

Ver el archivo `FIREBASE_EXAMPLE.ts` para ejemplos de uso.

#### Servicios disponibles en AngularFire

```typescript
import { Firestore } from '@angular/fire/firestore';
import { Auth } from '@angular/fire/auth';
import { Storage } from '@angular/fire/storage';
import { Functions } from '@angular/fire/functions';
```

## 🔥 Ejemplos rápidos

### Leer datos
```typescript
import { Firestore, collection, collectionData } from '@angular/fire/firestore';

constructor(private firestore: Firestore) {
  const itemsCollection = collection(this.firestore, 'items');
  collectionData(itemsCollection).subscribe(items => {
    console.log(items);
  });
}
```

### Escribir datos
```typescript
import { Firestore, doc, setDoc } from '@angular/fire/firestore';

async addData() {
  const docRef = doc(this.firestore, 'items', 'item1');
  await setDoc(docRef, { name: 'Item 1', value: 100 });
}
```

### Autenticación
```typescript
import { Auth, signInWithEmailAndPassword } from '@angular/fire/auth';

async login(email: string, password: string) {
  const userCredential = await signInWithEmailAndPassword(
    this.auth, 
    email, 
    password
  );
  console.log('Usuario:', userCredential.user);
}
```

## 📚 Recursos

- [AngularFire Documentation](https://github.com/angular/angularfire)
- [Firebase Documentation](https://firebase.google.com/docs)
- [Firestore Documentation](https://firebase.google.com/docs/firestore)
- [Firebase Authentication](https://firebase.google.com/docs/auth)

## ⚠️ Notas importantes

1. **Nunca** commits tus credenciales de Firebase en repositorios públicos
2. Usa variables de entorno para las credenciales en producción
3. Configura correctamente las reglas de seguridad antes de ir a producción
4. Habilita App Check para proteger tu backend de Firebase
