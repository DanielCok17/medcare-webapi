const mongoHost = process.env.AMBULANCE_API_MONGODB_HOST
const mongoPort = process.env.AMBULANCE_API_MONGODB_PORT

const mongoUser = process.env.AMBULANCE_API_MONGODB_USERNAME
const mongoPassword = process.env.AMBULANCE_API_MONGODB_PASSWORD

const database = process.env.AMBULANCE_API_MONGODB_DATABASE
const collections = process.env.AMBULANCE_API_MONGODB_COLLECTION.split(',')

const retrySeconds = parseInt(process.env.RETRY_CONNECTION_SECONDS || "6") || 6;

// try to connect to mongoDB until it is not available
let connection;
while (true) {
    try {
        connection = new Mongo(`mongodb://${mongoUser}:${mongoPassword}@${mongoHost}:${mongoPort}`);
        
        break;
    } catch (exception) {
        print(`Cannot connect to mongoDB: ${exception}`);
        print(`mongodb://${mongoUser}:${mongoPassword}@${mongoHost}:${mongoPort}`)
        print(`Will retry after ${retrySeconds} seconds`)
        sleep(retrySeconds * 1000);
    }
}

// if database and collections exist, exit with success - already initialized
const databases = connection.getDBNames();
if (databases.includes(database)) {
    const dbInstance = connection.getDB(database);
    const existingCollections = dbInstance.getCollectionNames();
    let allCollectionsExist = true;
    for (const collection of collections) {
        if (!existingCollections.includes(collection.split(':')[0])) {
            allCollectionsExist = false;
            break;
        }
    }
    if (allCollectionsExist) {
        print(`Collections '${collections.join(', ')}' already exist in database '${database}'`)
        process.exit(0);
    }
}

// initialize
const db = connection.getDB(database);
for (const collection of collections) {
    const collectionName = collection.split(':')[0];
    db.createCollection(collectionName);
    db[collectionName].createIndex({ "id": 1 });

    // Insert sample data specific to each collection if necessary
    // trigger commit for flux
    let sampleData = [];
    if (collectionName === "allergy_records") {
        sampleData = [
            {
                "id": "sample-allergy-record",
                "patientId": "sample-patient",
                "allergen": "Peanuts"
            }
        ];
    } else if (collectionName === "lab_results") {
        sampleData = [
            {
                "id": "sample-lab-result",
                "patientId": "sample-patient",
                "testType": "Blood Test",
                "result": "Normal"
            }
        ];
    } else if (collectionName === "medical_records") {
        sampleData = [
            {
                "id": "sample-medical-record",
                "patientId": "sample-patient",
                "condition": "Diabetes",
                "treatment": "Insulin",
                "history": "Long-term"
            }
        ];
    } else if (collectionName === "vaccination_records") {
        sampleData = [
            {
                "id": "sample-vaccination-record",
                "patientId": "sample-patient",
                "vaccine": "COVID-19",
                "date": "2024-01-15"
            }
        ];
    }

    let result = db[collectionName].insertMany(sampleData);
    if (result.writeError) {
        console.error(result)
        print(`Error when writing the data to collection '${collectionName}': ${result.errmsg}`)
    }
}

// exit with success
process.exit(0);
