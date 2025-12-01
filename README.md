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

## 2. Motor de Evaluación de Riesgo (IA Mock)

Para la **Opción 1 — Gestor de solicitudes de crédito y pre-evaluación de riesgo**, implementé un **motor de scoring crediticio** dentro del paquete `risk/`.  
Este componente actúa como una **IA mock**, simulando el comportamiento de un sistema de crédito real mediante reglas de negocio claras, explicables y desacopladas.

---

###  ¿Qué hace este módulo?

El archivo define la función principal `EvaluateCreditRisk`, encargada de:

1. **Cargar toda la información necesaria desde la BD**
   - Solicitud de crédito (`CreditRequest`)
   - Cliente asociado (`Customer`)
   - Historial de créditos previos
   - Activos asociados específicamente a esa solicitud

2. **Calcular un puntaje numérico de riesgo (0–100)**
   Utiliza la función `calculateScore`, que evalúa factores clave:
   - **Relación cuota / ingreso:** Mientras menor sea, menor riesgo.
   - **Activos registrados:** Se calcula la relación `valor_activos / monto_solicitado`.
   - **Existencia de vivienda como respaldo:** Mejora la estabilidad financiera.
   - **Historial crediticio:**  
     - Créditos aprobados → aumenta el score  
     - Créditos rechazados o muchas solicitudes → lo disminuyen  
   - **Tipo de producto:**  
     - Vivienda/Hipotecario → bajo riesgo  
     - Libre inversión/consumo → mayor riesgo  

   El score inicia en **50** y se ajusta según reglas.  
   Finalmente se limita al rango **0 a 100**.

3. **Clasificar el riesgo**
   Con `riskCategory(score)` se obtiene:
   - **LOW / Bajo** (≥ 80)
   - **MEDIUM / Medio** (≥ 55)
   - **HIGH / Alto** (< 55)

4. **Sugerir una decisión crediticia**
   `recommendationFromScore(score)` retorna:
   - `APROBAR`
   - `DEJAR EN ESTUDIO / APROBAR CON CONDICIONES`
   - `NO APROBAR`

5. **Generar explicación en lenguaje natural**
   Con `buildExplanation` se construye una salida tipo *IA explicable* que incluye:
   - Puntaje final.  
   - Categoría de riesgo.  
   - Recomendación.  
   - Razones justificadas basadas en datos reales.  
   - Mejoras sugeridas (por ejemplo, registrar activos, reducir monto, etc).

6. **Guardar resultados en la solicitud**
   Se guardan:
   - `RiskScore`
   - `RiskCategory`
   - `RiskExplanation`

   Cuando se modifica algun detalle del credito el sistema automaticamente genera un nuevo reporte financiero

---

###  ¿Por qué esto es una IA Mock?

- Simula el comportamiento de un modelo de scoring real.
- Produce una explicación textual detallada sobre la solicitud de credito.
- Está desacoplado, por lo que puede ser reemplazada en el futuro por un modelo real (ML/LLM).
- El diseño facilita conectarlo a un endpoint externo de IA si el proyecto lo requiere.

---

### Relación con el punto entregado


Se implementó la **Opción N°1** propuesta en la prueba técnica:

> **Sistema gestor de solicitudes de crédito con pre-evaluación de riesgo y explicación en lenguaje natural.**

El módulo `risk/` representa el **núcleo inteligente del sistema**, encargado de calcular el puntaje de riesgo, clasificarlo en un rango (Bajo, Medio o Alto) y generar una explicación detallada que justifica la decisión crediticia.

Además, el modelo funciona bajo un principio fundamental:

### Entre mayor sea el puntaje (score), mayor es la probabilidad de aprobación del crédito.

Un puntaje alto indica:
- Mejor relación cuota/ingreso  
- Mayor respaldo en activos  
- Mejor historial crediticio  
- Menor riesgo financiero  

Y, por tanto, una **mayor posibilidad de que el crédito sea aprobado**.

Con esto, la funcionalidad cumple completamente el objetivo de la Opción N°1, entregando un sistema que evalúa solicitudes y proporciona una explicación clara, transparente y útil para toma de decisiones.

---

### Resultado final

El motor permite que cada solicitud:
- Sea evaluada automáticamente,
- Obtenga un puntaje cuantitativo,
- Obtenga una categoría de riesgo,
- Reciba una explicación justificable y legible por un analista,
- Y quede almacenada con su evaluación.

Cada vez que se realiza un cambio en la información del credito **se genera una nueva evaluación** de forma automática, garantizando información actualizada y confiable.

---

## **3. Instrucciones para levantar el entorno con Docker**

Antes de iniciar el proyecto, asegúrate de cumplir con los siguientes requisitos:

### Requisitos Previos

#
Docker
Debes tener instalado en tu computador:
 - Docker Desktop

Docker es indispensable para ejecutar el proyecto con contenedores.

### Node.js
(Solo si vas a utilizar el frontend fuera de Docker)
No es necesario tener Node.js para ejecutar el proyecto con Docker.
Solo instálalo si deseas trabajar localmente en el frontend:
- Node.js v20.12.2 o superior
- npm o yarn

 Puertos Necesarios
 Asegúrate de que los siguientes puertos estén libres:

 ### Puertos utilizados


| Servicio            |   Puerto    | Descripción            |
|---------------------|:-----------:|------------------------|
| PostgreSQL          |    5435     | Base de datos          |
| Backend (Go)        | 5000 / 4000 | API REST del backend   |
| Frontend (Next.js)  |    3000     | Aplicación web         |


 Si alguno de estos puertos está ocupado, Docker no podrá iniciar correctamente los servicios.

#### PostgreSQL
Debes contar con **PostgreSQL instalado localmente**  
o con **credenciales de acceso remoto a una base de datos PostgreSQL**:
- Host
- Puerto
- Nombre de base de datos
- Usuario
- Contraseña

> Por defecto, el proyecto usa PostgreSQL expuesto en el puerto **5435**.  
> Si utilizas una instancia externa, ajusta las variables de entorno para que apunten a tu servidor con PostgreSQL.


### **Paso 1 — Clonar el repositorio**
```bash
git clone -b main https://github.com/JhonCamargo53/credit-risk-go-next-postgres.git

cd credit-risk-go-next-postgres
```
###  **Paso 2 — Establecer las variables de entorno**

El proyecto cuenta con archivos `.env` tanto para **desarrollo** como para **producción** en el frontend y backend, puedes modificarlos según tus necesidades.  

Para facilitar el despliegue, el repositorio **ya incluye** las variables configuradas.

---

### **Variables de entorno — Backend**
Asegúrate de:

- Cambiar usuario, contraseña y nombre de base de datos a los de tu entorno *(necesario obligatoriamente en local , si lo haces con docker el mismo crea la base de datos)*.
- Verificar que las URLs del backend coincidan con los puertos que estás usando (local o Docker).

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

### **Paso 3 — Ejecutar el proyecto**


**Docker:** Una vez configuradas las variables de entorno del backend y frontend, el siguiente paso es levantar toda la infraestructura utilizando Docker Compose.

Ejecuta el siguiente comando:

```bash
docker compose up --build
```
---

**Local:** Ya teniendo la variables configuradas se procede a correr el proyecto de manera loca, para el frotend escribe en cosola `npm i` para instalar las dependencias una vez termine escriba `npm run dev` y ya estara funcionando el proyecto NextJs.

Para el backend es necesario un unico comando `air` al ejecutar saldra en consola el estado del servidor.

## **Paso 4 — Comprobación de funcionamiento**

Una vez ejecutados los pasos anteriores y siempre que no se hayan modificado las variables de entorno, tendrás correctamente levantados los tres servicios del proyecto:

- **Backend (Go):** `http://localhost:5000`  
- **Frontend (Next.js):**` http://localhost:3000`
- **PostgreSQL:** `localhost:5435`  

---

###  Verificación del Backend 

Para comprobar el backend de docker está funcionando correctamente accediendo en tu navegador a: `http://localhost:5000/health`. Si el servicio está activo, verás una respuesta indicando que el servidor está operativo.

Tambien puedes acceder a la documentación utilizando el siguiente enlace `http://localhost:5000/swagger/index.html`

Si deseas comprobar el backend local utiliza el el siguiente enlace `http://localhost:4000/health`

---

### Verificación del Frontend

Para verificar que el frontend de docker está funcionando, abre tu navegador y accede a: `http://localhost:3000`. Ahí podrás visualizar la interfaz principal de la aplicación y comenzar a interactuar con el sistema.

*Nota: El frontend corre en el mismo puerto estando en local o desde docker.*

---

## **Acceso a la Plataforma**

| Servicio | URL | Descripción |
|---------|-----|-------------|
| **Frontend (Producción)** | https://risk-management.alphacodexs.com | Interfaz web para gestión de clientes, solicitudes y análisis de riesgo. |
| **Backend (Producción)** |  https://risk-management-backend.alphacodexs.com | API REST del sistema, motor de evaluación y autenticación. |

---

> Puedes **crear usuarios de prueba**, registrar clientes, agregar activos, generar solicitudes, ejecutar evaluaciones de riesgo y visualizar el flujo completo de funcionamiento del sistema.

---
### **¿Qué puedes probar en producción?**
- Inicio de sesión.
- Gestionar usuarios. 
- Gestión de clientes y sus activos.  
- Creación de solicitudes de crédito.  
- Motor de evaluación de riesgo (IA Mock).  
- Reportes financieros generados automáticamente.  

---

> **Si no dispone de credenciales para acceder al entorno de producción, puede solicitarlas para realizar las pruebas correspondientes.**

## 📁 Documentación del Proyecto

En la carpeta **`/docs`** se encuentran los diagramas principales del sistema.  
A continuación puedes verlos directamente en vista previa:

---

### Diagrama de Arquitectura
[![Diagrama de Arquitectura](https://github.com/JhonCamargo53/credit-risk-go-next-postgres/blob/main/docs/Diagrama%20de%20Arquitectura.svg)](https://github.com/JhonCamargo53/credit-risk-go-next-postgres/blob/main/docs/Diagrama%20de%20Arquitectura.svg)

---

### Diagrama de la Base de Datos
[![Diagrama de la Base de Datos](https://github.com/JhonCamargo53/credit-risk-go-next-postgres/blob/main/docs/Diagrama%20de%20la%20base%20de%20datos.png)](https://github.com/JhonCamargo53/credit-risk-go-next-postgres/blob/main/docs/Diagrama%20de%20la%20base%20de%20datos.png)

---