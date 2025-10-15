/**
 * Тестовый скрипт для проверки подключения к MongoDB
 */

// Load environment variables
require('dotenv').config();

const { MongoClient } = require('mongodb');

async function test() {
  const uri = process.env.MONGODB_URI || 'mongodb://localhost:27017/fastspot';
  const client = new MongoClient(uri);
  
  try {
    console.log('🔌 Connecting to MongoDB...');
    console.log(`📍 URI: ${uri.replace(/\/\/.*@/, '//***@')}`); // hide credentials for security reasons
    
    await client.connect();
    console.log('✅ Connection successful!\n');
    
    const db = client.db();
    const collections = await db.listCollections().toArray();
    
    console.log(`📊 Database: ${db.databaseName}`);
    console.log(`📦 Found collections: ${collections.length}\n`);
    
    if (collections.length > 0) {
      console.log('📋 List of collections:');
      
      for (const collection of collections) {
        const count = await db.collection(collection.name).countDocuments();
        console.log(`   ${collection.name.padEnd(20)} - ${count} documents`);
      }
    } else {
      console.log('⚠️  Database is empty. Run: npm run seed');
    }
    
    console.log('\n🎉 Test completed successfully!');
    
  } catch (error) {
    console.error('\n❌ Error connecting:', error.message);
    console.error('\n💡 Solutions:');
    console.error('   1. Make sure MongoDB is running');
    console.error('   2. Check MONGODB_URI in .env file');
    console.error('   3. Check login/password (if using authentication)');
    process.exit(1);
  } finally {
    await client.close();
    console.log('🔌 Connection closed\n');
  }
}

test();

