# Chat-PSC

## Descripción

Este proyecto implementa un servidor de chat utilizando ZeroMQ y sockets de tipo ROUTER y DEALER para la asignatura de PSC. El servidor permite a los clientes conectarse, enviar mensajes y recibir mensajes específicos dirigidos a ellos. Los clientes pueden ver los mensajes recibidos y enviar nuevos mensajes a otros usuarios conectados.

## Funcionamiento

1. **Servidor**:
   - El servidor utiliza un socket ROUTER para recibir y enviar mensajes.
   - Cuando un cliente se conecta, el servidor almacena su nombre de usuario y le envía un mensaje de bienvenida.
   - El servidor almacena los mensajes y los reenvía al destinatario especificado en el mensaje, permitiendo que un usuario se pueda desconectar y volver a conectar sin perder los mensajes recibidos.

2. **Cliente**:
   - El cliente utiliza un socket DEALER para comunicarse con el servidor.
   - El cliente se conecta al servidor con un nombre de usuario.
   - El cliente puede enviar mensajes a otros usuarios y ver los mensajes recibidos.

## Limitaciones

- **Persistencia**: Los mensajes no se almacenan de manera persistente. Si el servidor se reinicia, se perderán todos los mensajes almacenados en memoria.
- **Seguridad**: No se implementa ninguna medida de seguridad, como cifrado o autenticación. En este proyecto se implementa únicamente la lógica del chat.
- **Escalabilidad**: El servidor actual está diseñado para funcionar en un solo nodo y puede no escalar bien con un gran número de clientes.

## Requisitos

- Go
- ZeroMQ (versión 4.x)

## Ejecución del chat

1. Clonar el repositorio:
    ```sh
    git clone --depth 1 --branch v1.0.0 https://github.com/JorgeSanchezL/Chat-PSC
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

4. Ejecutar el servidor
    ```sh
    ./server
    ```

5. En terminales diferentes a la del servidor, ejecutar tantos clientes como queramos
    ```sh
    ./client <Nombre de usuario>
    ```

6. Una vez conectado, puedes usar el siguiente comando:
  - `s`: Enviar un nuevo mensaje. Se te pedirá que ingreses el destinatario y el contenido del mensaje.