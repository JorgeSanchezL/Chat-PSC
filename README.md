# Chat-PSC

## Descripción

Este proyecto implementa un servidor de chat utilizando NATS para la asignatura de PSC. Se divide en dos partes (cliente y servidor), pero puede ser ejecutado únicamente mediante el cliente, sin la presencia de ningún servidor. Los clientes pueden enviar y recibir mensajes utilizando NATS, mientras que el servidor se encarga de gestionar los usarios conectados y guardar los mensajes de forma temporal, hasta que el servidor se reinicia.

## Funcionamiento

1. **Servidor**:
   - El servidor utiliza dos subscripciones, una para gestionar la conexión de usuarios nuevos y otra para sincronizar las listas de usuarios locales de los clientes con la del servidor.
   - Cuando un cliente se conecta, el servidor almacena su nombre de usuario.
   - El servidor almacena los mensajes de cada usuario y los publica en el tema correspondiente cuando el usuario vuelve a conectarse, permitiendo que un usuario se pueda desconectar sin perder los mensajes.

2. **Cliente**:
   - El cliente utiliza una subscripcion wildcard para poder recibir mensajes de dos temas diferentes (old y new), dependiendo si son los mensajes recibidos en tiempo real o los que ya existian antes de su última conexión.
   - El cliente se conecta al servidor con un nombre de usuario.
   - El cliente puede enviar mensajes a otros usuarios y ver los mensajes recibidos.

## Limitaciones

- **Persistencia**: Los mensajes no se almacenan de manera persistente. Si el servidor se reinicia, se perderán todos los mensajes almacenados en memoria. Para evitar este problema, teniendo en cuenta que estamos utilizando NATS, podríamos activar JetStream en la configuración de NATS, eliminar el servidor, y cambiar las publicaciones y subscripciones por almacenamientos en JetStream y observadores sobre los objetos de JetStream.
- **Seguridad**: No se implementa ninguna medida de seguridad, como cifrado o autenticación. En este proyecto se implementa únicamente la lógica del chat. De nuevo, con la persistencia de JetStream y añadiendo algo de lógica al servidor podríamos implementar un sistema de inicio de sesión para garantizar la seguridad.
- **Escalabilidad**: El servidor actual está diseñado para funcionar en un solo nodo y puede no escalar bien con un gran número de clientes. No obstante, al permitir la ejecucion de los clientes sin ningún servidor, mejora bastante la escalabilidad. Por el contrario, si se lleva a cabo la solución mencionada de añadir JetStream para implementar la persistencia, dado que se trata de un almacenamiento de clave-valor cuya clave sería el usuario, el tiempo de procesamiento de un mensaje aumentaría mucho al aumentar el número de mensajes almacenados por usuario, ya que el valor contendría todos los mensajes anteriores.

## Requisitos

- Go
- NATS

## Variables de entorno
- NATS_URL: Usada para inicializar NATS tanto en el servidor como en el cliente. Si no tiene ningún valor asignado, se utiliza la URL por defecto ("nats://127.0.0.1:4222")

## Ejecución del chat

1. Clonar el repositorio:
    ```sh
    git clone --depth 1 --branch v2.0.0 https://github.com/JorgeSanchezL/Chat-PSC
    cd Chat-PSC
    ```

2. Compilar el servidor:
    ```sh
    cd chat-server
    go build
    mv server ..
    cd ..
    ```

3. Compilar el cliente:
    ```sh
    cd chat-client
    go build
    mv client ..
    cd ..
    ```

4. Ejecutar NATS. Con el siguiente comando podemos ejecutarlo dentro de un contenedor de Docker:
    ```sh
    docker run --name natssrv --rm -p 4222:4222 nats
    ```

5. Ejecutar el servidor (opcional):
    ```sh
    ./server
    ```

6. En terminales diferentes a la del servidor, ejecutar tantos clientes como queramos:
    ```sh
    ./client <Nombre de usuario>
    ```

7. Una vez conectado, puedes usar el siguiente comando:
  - `s`: Enviar un nuevo mensaje. Se te pedirá que ingreses el destinatario y el contenido del mensaje.