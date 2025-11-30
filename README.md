# **Gesto de Solicitudes de Credito**
### _Backend en Go · Frontend en Next.js + TailwindCSS + React Context + Zustand · PostgreSQL · Motor de Evaluación Simulado · Tokenización Blockchain_

---

##  **1. Problema elegido y justificación**

Elegí esta opción porque es la que más se relaciona con mi interés personal por **comprender cómo funciona el historial crediticio de los clientes** y cómo las entidades financieras **procesan esta información internamente** para tomar decisiones.  

Me resulta interesante profundizar en los criterios que influyen en la evaluación crediticia, como:

- **Ingresos**
- **Monto solicitado**
- **Plazo**
- **Comportamiento financiero**
- **Relación monto/ingreso**
- **Historial del cliente**

Todos estos factores se relacionan entre sí para producir un **análisis de riesgo personalizado**.  
Esta opción me permite explorar precisamente esa lógica, investigando qué variables tienen mayor impacto al momento de aprobar un crédito y cómo los bancos convierten datos básicos de un cliente en un **reporte financiero de estado**, que luego se utiliza para aprobar, rechazar o solicitar análisis adicional sobre una solicitud.

---

### **Reto técnico principal**

El principal reto técnico de esta opción será la **construcción del mock que simula el comportamiento de una IA**, encargado de:

- Procesar los datos del cliente y de la solicitud  
- Generar un puntaje o categoría de riesgo (**bajo, medio o alto**)  
- Producir un **reporte explicativo en lenguaje natural** que justifique la decisión  

Para resolverlo, es necesario:

- Definir un conjunto **consistente de reglas o pesos**
- Analizar la relación entre variables como **monto/ingreso**, tipo de producto y plazo  
- Convertir este razonamiento en una **salida clara, coherente y replicable**

Además, su implementación requiere **desacoplar la lógica mediante un puerto de dominio**, siguiendo la **arquitectura hexagonal**, permitiendo sustituir este mock posteriormente por un modelo real o un servicio externo sin afectar la estructura del sistema.

---

### **Propuesta Blockchain**

Como valor agregado opcional, propongo utilizar **blockchain para tokenizar el reporte financiero generado**.  

La idea consiste en emitir un **token único e inmutable** asociado al resultado del análisis de riesgo, garantizando:

- Que el reporte no pueda ser alterado sin dejar evidencia  
- **Integridad** y **trazabilidad** de cada documento  
- Auditorías más confiables  
- Que cada reporte se convierta en un **activo verificable** dentro de una red blockchain  

Esto mejora la **transparencia**, fortalece la **confianza del sistema** y permite que las decisiones queden registradas de manera permanente e inmutable.

---

## **2. Instrucciones para levantar el entorno con Docker**

El proyecto incluye tres servicios principales:

- **Backend (Go)**
- **Frontend (Next.js + TailwindCSS + Context + Zustand)**
- **PostgreSQL**
---

### **Paso 1 — Clonar el repositorio**
```bash
git clone **AUN NO LO CREO**
cd REPO
```
###  **Paso 2 — Establecer las variables de entorno**

El proyecto cuenta con archivos `.env` tanto para **desarrollo** como para **producción** en el frontend y backend.  
Puedes modificarlos según tus necesidades.  
Para facilitar el despliegue, el repositorio **ya incluye** las variables configuradas.

---

### **Variables de entorno — Backend**

### Archivo: `.env.development`
```env
ENV=development
DATABASE_URL=host=localhost user=admin password=20Acc3ss25 dbname=credit port=5435 sslmode=disable
PORT=4000
JWT_SECRET_KEY=CE8PvUkf92YJdOXKXm8Aw43BKTOgua...
```

#### Archivo: `.env.production`
```bash
ENV=production
DATABASE_URL=host=db user=admin password=20Acc3ss25 dbname=credit port=5432 sslmode=disable
PORT=5000
JWT_SECRET_KEY=J0vadRYJdOXKX2YJdOXKJXm8HAw43BK...
```

### **Variables de entorno — Frontend**

#### Archivo: `.env.development`
```env
NEXT_PUBLIC_NODE_ENV=development
NEXT_PUBLIC_DEVELOP_BASE_URL=http://localhost:4000/
NEXT_PUBLIC_PRODUCTION_BASE_URL=http://localhost:4000/
NEXT_PUBLIC_JWT_COOKIE_NAME=credit-management-token
```

### Archivo: `.env.production`
```bash
NEXT_PUBLIC_NODE_ENV=production
NEXT_PUBLIC_DEVELOP_BASE_URL=http://localhost:5000/
NEXT_PUBLIC_PRODUCTION_BASE_URL=http://localhost:5000/
NEXT_PUBLIC_JWT_COOKIE_NAME=credit-management-token
```

---

### **Paso 3 — Ejecutar el proyecto con Docker**

Una vez configuradas las variables de entorno del backend y frontend, el siguiente paso es levantar toda la infraestructura utilizando Docker Compose.

Ejecuta el siguiente comando:

```bash
docker compose up --build
```
---

## **Paso 4 — Comprobación de funcionamiento**

Una vez ejecutados los pasos anteriores y siempre que no se hayan modificado las variables de entorno, tendrás correctamente levantados los tres servicios del proyecto:

- **Backend (Go):** `http://localhost:5000`  
- **Frontend (Next.js):**` http://localhost:3000`
- **PostgreSQL:** `localhost:5435`  

---

###  Verificación del Backend

Puedes comprobar que el backend está funcionando correctamente accediendo en tu navegador a: `http://localhost:5000/health`. Si el servicio está activo, verás una respuesta indicando que el servidor está operativo.

Tambien puedes acceder a la documentación utilizando el siguiente enlace `http://localhost:5000/swagger/index.html`

---

### Verificación del Frontend

Para verificar que el frontend está funcionando, abre tu navegador y accede a: `http://localhost:3000`. Ahí podrás visualizar la interfaz principal de la aplicación y comenzar a interactuar con el sistema.