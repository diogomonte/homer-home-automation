#include <ESP8266WiFi.h>
#include <PubSubClient.h>

const char* ssid = "TN_24GHz_F30D01";
const char* password = "WJTNVULYEN";
const char* mqtt_server = "192.168.10.188";

const char* action_topic = "homer/workstation_screenlight/action";
const char* event_topic = "homer/workstation_screenlight/event";

WiFiClient espClient;
PubSubClient client(espClient);

void connect_to_wifi() {

  delay(10);
  Serial.println();
  Serial.print("Connecting to ");
  Serial.println(ssid);

  WiFi.mode(WIFI_STA);
  WiFi.begin(ssid, password);

  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.println("trying to connect...");
  }

  randomSeed(micros());

  Serial.println("");
  Serial.println("WiFi connected");
  Serial.println("IP address: ");
  Serial.println(WiFi.localIP());
}

void callback(char* topic, byte* payload, unsigned int length) {
  Serial.print("Message arrived [");
  Serial.print(topic);
  Serial.print("] ");
  for (int i = 0; i < length; i++) {
    Serial.print((char)payload[i]);
  }
  Serial.println();

  // Switch on the LED if an 1 was received as first character
  if ((char)payload[0] == '1') {
    digitalWrite(BUILTIN_LED, LOW);   // Turn the LED on (Note that LOW is the voltage level
    // but actually the LED is on; this is because
    // it is active low on the ESP-01)
  } else {
    digitalWrite(BUILTIN_LED, HIGH);  // Turn the LED off by making the voltage HIGH
  }

}

void reconnect() {
  while (!client.connected()) {
    Serial.print("Attempting MQTT connection...");
    // Create a random client ID
    String clientId = "workstation-light" + String(random(0xffff), HEX);
    if (client.connect(clientId.c_str())) {
      Serial.println("connected to mqtt server");
      
      String message_id = String(random(0xffff), HEX);
      String message = create_mqtt_message(message_id, "connected", WiFi.localIP().toString());
      client.publish(event_topic, message.c_str());
      client.subscribe(action_topic);
    } else {
      Serial.print("failed, rc=");
      Serial.print(client.state());
      Serial.println(" try again in 5 seconds");
      // Wait 5 seconds before retrying
      delay(5000);
    }
  }
}

String create_mqtt_message(String messageId, String event_key, String event) {
  String message = "";
  message.concat("{");
  message.concat("'header': {'message_id': '" + messageId + "'},");
  message.concat("'body': {'" + event_key + "':'" + event + "'}");
  message.concat("}");
  return message;
}

void setup() {
  pinMode(BUILTIN_LED, OUTPUT);     // Initialize the BUILTIN_LED pin as an output
  Serial.begin(115200);
  connect_to_wifi();
  client.setServer(mqtt_server, 1883);
  client.setCallback(callback);
}

void loop() {
  if (!client.connected()) {
    reconnect();
  }
  client.loop();

  String message_id = String(random(0xffff), HEX);
  String message = create_mqtt_message(message_id, "ping", WiFi.localIP().toString());
  client.publish(event_topic, message.c_str());
          
  delay(5000);
}
