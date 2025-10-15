/**
 * FastSpot - MongoDB Seed Script
 **/

// Load environment variables
require('dotenv').config();

const { MongoClient, ObjectId } = require('mongodb');
const bcrypt = require('bcryptjs');

// Configuration
const MONGODB_URI = process.env.MONGODB_URI || 'mongodb://localhost:27017/fastspot';
const DB_NAME = 'fastspot';

async function seed() {
  const client = new MongoClient(MONGODB_URI);
  
  try {
    console.log('🔌 Connecting to MongoDB...');
    await client.connect();
    const db = client.db(DB_NAME);
    
    console.log('🗑️  Clearing existing data...');
    await Promise.all([
      db.collection('users').deleteMany({}),
      db.collection('categories').deleteMany({}),
      db.collection('products').deleteMany({}),
      db.collection('promotions').deleteMany({}),
      db.collection('carts').deleteMany({}),
      db.collection('orders').deleteMany({}),
      db.collection('mood_questions').deleteMany({}),
      db.collection('mood_rules').deleteMany({}),
      db.collection('ai_sessions').deleteMany({})
    ]);

    // ==================== USERS ====================
    console.log('👥 Creating users...');
    const adminPassword = await bcrypt.hash('Admin123!', 10);
    
    const usersResult = await db.collection('users').insertMany([
      {
        role: 'admin',
        name: 'Admin',
        email: 'admin@local',
        phone: '+1234567890',
        passwordHash: adminPassword,
        createdAt: new Date()
      }
    ]);
    console.log(`✅ Users created: ${usersResult.insertedCount}`);

    // ==================== CATEGORIES ====================
    console.log('📂 Creating categories...');
    const categoriesResult = await db.collection('categories').insertMany([
      {
        _id: new ObjectId(),
        name: 'Burgers',
        slug: 'burgers',
        image: 'https://images.unsplash.com/photo-1568901346375-23c9450c58cd?w=400',
        isActive: true
      },
      {
        _id: new ObjectId(),
        name: 'Drinks',
        slug: 'drinks',
        image: 'https://images.unsplash.com/photo-1437418747212-8d9709afab22?w=400',
        isActive: true
      },
      {
        _id: new ObjectId(),
        name: 'Desserts',
        slug: 'desserts',
        image: 'https://images.unsplash.com/photo-1551024506-0bccd828d307?w=400',
        isActive: true
      },
      {
        _id: new ObjectId(),
        name: 'Snacks',
        slug: 'snacks',
        image: 'https://images.unsplash.com/photo-1541592106381-b31e9677c0e5?w=400',
        isActive: true
      }
    ]);
    
    const categories = await db.collection('categories').find({}).toArray();
    const [burgers, drinks, desserts, snacks] = categories;
    console.log(`✅ Categories created: ${categoriesResult.insertedCount}`);

    // ==================== PRODUCTS ====================
    console.log('🍔 Creating products...');
    const productsResult = await db.collection('products').insertMany([
      // Burgers
      {
        _id: new ObjectId(),
        categoryId: burgers._id,
        name: 'Classic Burger',
        slug: 'classic-burger',
        description: 'Juicy beef patty, fresh vegetables, cheddar cheese and signature sauce',
        priceUSD: 5.99,
        image: 'https://images.unsplash.com/photo-1568901346375-23c9450c58cd?w=600',
        isActive: true,
        ingredients: [
          { key: 'pickles', label: 'Marinated cucumbers', defaultIncluded: true },
          { key: 'onions', label: 'Onions', defaultIncluded: true },
          { key: 'tomatoes', label: 'Tomatoes', defaultIncluded: true },
          { key: 'lettuce', label: 'Lettuce', defaultIncluded: true }
        ],
        options: [
          {
            key: 'size',
            label: 'Size',
            type: 'single',
            choices: [
              { value: 'regular', label: 'Regular', extraPriceUSD: 0 },
              { value: 'large', label: 'Large', extraPriceUSD: 2.00 }
            ]
          },
          {
            key: 'sauce',
            label: 'Sauce',
            type: 'single',
            choices: [
              { value: 'classic', label: 'Classic', extraPriceUSD: 0 },
              { value: 'bbq', label: 'BBQ', extraPriceUSD: 0.50 },
              { value: 'spicy', label: 'Spicy', extraPriceUSD: 0.50 }
            ]
          }
        ],
        tags: ['comfort-food', 'popular', 'savory']
      },
      {
        _id: new ObjectId(),
        categoryId: burgers._id,
        name: 'Deluxe Cheeseburger',
        slug: 'deluxe-cheeseburger',
        description: 'Double patty, double cheese, bacon and signature sauce',
        priceUSD: 8.49,
        image: 'https://images.unsplash.com/photo-1572802419224-296b0aeee0d9?w=600',
        isActive: true,
        ingredients: [
          { key: 'pickles', label: 'Marinated cucumbers', defaultIncluded: true },
          { key: 'onions', label: 'Onions', defaultIncluded: true },
          { key: 'bacon', label: 'Bacon', defaultIncluded: true }
        ],
        options: [
          {
            key: 'cheese',
            label: 'Additional cheese',
            type: 'single',
            choices: [
              { value: 'no', label: 'No', extraPriceUSD: 0 },
              { value: 'yes', label: 'Yes (+$1)', extraPriceUSD: 1.00 }
            ]
          }
        ],
        tags: ['comfort-food', 'indulgent', 'savory', 'popular']
      },
      {
        _id: new ObjectId(),
        categoryId: burgers._id,
        name: 'Veggie Burger',
        slug: 'veggie-burger',
        description: 'Veggie patty, avocados, fresh vegetables',
        priceUSD: 6.99,
        image: 'https://images.unsplash.com/photo-1520072959219-c595dc870360?w=600',
        isActive: true,
        ingredients: [
          { key: 'avocado', label: 'Avocado', defaultIncluded: true },
          { key: 'tomatoes', label: 'Tomatoes', defaultIncluded: true },
          { key: 'lettuce', label: 'Lettuce', defaultIncluded: true }
        ],
        options: [],
        tags: ['vegetarian', 'healthy', 'light']
      },

      // Drinks
      {
        _id: new ObjectId(),
        categoryId: drinks._id,
        name: 'Coca-Cola',
        slug: 'coca-cola',
        description: 'Refreshing carbonated drink',
        priceUSD: 2.49,
        image: 'https://images.unsplash.com/photo-1554866585-cd94860890b7?w=600',
        isActive: true,
        ingredients: [],
        options: [
          {
            key: 'size',
            label: 'Size',
            type: 'single',
            choices: [
              { value: 'small', label: 'Small (0.33L)', extraPriceUSD: 0 },
              { value: 'medium', label: 'Medium (0.5L)', extraPriceUSD: 0.50 },
              { value: 'large', label: 'Large (0.75L)', extraPriceUSD: 1.00 }
            ]
          },
          {
            key: 'ice',
            label: 'Ice',
            type: 'single',
            choices: [
              { value: 'regular', label: 'Regular', extraPriceUSD: 0 },
              { value: 'no-ice', label: 'No ice', extraPriceUSD: 0 },
              { value: 'extra-ice', label: 'Extra ice', extraPriceUSD: 0 }
            ]
          }
        ],
        tags: ['refreshing', 'sweet']
      },
      {
        _id: new ObjectId(),
        categoryId: drinks._id,
        name: 'Orange Juice',
        slug: 'orange-juice',
        description: '100% natural fresh juice',
        priceUSD: 3.99,
        image: 'https://images.unsplash.com/photo-1600271886742-f049cd451bba?w=600',
        isActive: true,
        ingredients: [],
        options: [
          {
            key: 'size',
            label: 'Size',
            type: 'single',
            choices: [
              { value: 'small', label: 'Small (0.25L)', extraPriceUSD: 0 },
              { value: 'large', label: 'Large (0.5L)', extraPriceUSD: 1.50 }
            ]
          }
        ],
        tags: ['healthy', 'refreshing', 'energizing']
      },
      {
        _id: new ObjectId(),
        categoryId: drinks._id,
        name: 'Milkshake',
        slug: 'milkshake',
        description: 'Thick creamy shake with different flavors',
        priceUSD: 4.99,
        image: 'https://images.unsplash.com/photo-1572490122747-3968b75cc699?w=600',
        isActive: true,
        ingredients: [],
        options: [
          {
            key: 'flavor',
            label: 'Flavor',
            type: 'single',
            choices: [
              { value: 'vanilla', label: 'Vanilla', extraPriceUSD: 0 },
              { value: 'chocolate', label: 'Chocolate', extraPriceUSD: 0 },
              { value: 'strawberry', label: 'Strawberry', extraPriceUSD: 0 }
            ]
          }
        ],
        tags: ['sweet', 'indulgent', 'comfort-food']
      },

      // Desserts
      {
        _id: new ObjectId(),
        categoryId: desserts._id,
        name: 'Chocolate Brownie',
        slug: 'chocolate-brownie',
        description: 'Warm chocolate brownie with vanilla ice cream',
        priceUSD: 3.99,
        image: 'https://images.unsplash.com/photo-1607920591413-4ec007e70023?w=600',
        isActive: true,
        ingredients: [],
        options: [
          {
            key: 'ice-cream',
            label: 'Ice cream',
            type: 'single',
            choices: [
              { value: 'yes', label: 'With ice cream', extraPriceUSD: 0 },
              { value: 'no', label: 'Without ice cream', extraPriceUSD: 0 }
            ]
          }
        ],
        tags: ['sweet', 'indulgent', 'comfort-food']
      },
      {
        _id: new ObjectId(),
        categoryId: desserts._id,
        name: 'Apple Pie',
        slug: 'apple-pie',
        description: 'Classic American apple pie',
        priceUSD: 2.99,
        image: 'https://images.unsplash.com/photo-1535920527002-b35e96722eb9?w=600',
        isActive: true,
        ingredients: [],
        options: [],
        tags: ['sweet', 'comfort-food', 'classic']
      },

      // Snacks
      {
        _id: new ObjectId(),
        categoryId: snacks._id,
        name: 'French Fries',
        slug: 'french-fries',
        description: 'Crispy golden french fries',
        priceUSD: 2.99,
        image: 'https://images.unsplash.com/photo-1576107232684-1279f390859f?w=600',
        isActive: true,
        ingredients: [],
        options: [
          {
            key: 'size',
            label: 'Size',
            type: 'single',
            choices: [
              { value: 'small', label: 'Small', extraPriceUSD: 0 },
              { value: 'medium', label: 'Medium', extraPriceUSD: 1.00 },
              { value: 'large', label: 'Large', extraPriceUSD: 2.00 }
            ]
          }
        ],
        tags: ['savory', 'popular', 'comfort-food']
      },
      {
        _id: new ObjectId(),
        categoryId: snacks._id,
        name: 'Chicken Nuggets',
        slug: 'chicken-nuggets',
        description: 'Crispy chicken nuggets (6 pieces)',
        priceUSD: 4.99,
        image: 'https://images.unsplash.com/photo-1562967914-608f82629710?w=600',
        isActive: true,
        ingredients: [],
        options: [
          {
            key: 'sauce',
            label: 'Sauce',
            type: 'multiple',
            choices: [
              { value: 'bbq', label: 'BBQ', extraPriceUSD: 0 },
              { value: 'honey-mustard', label: 'Honey-mustard', extraPriceUSD: 0 },
              { value: 'ranch', label: 'Ranch', extraPriceUSD: 0 }
            ]
          }
        ],
        tags: ['savory', 'popular', 'comfort-food']
      }
    ]);
    
    const products = await db.collection('products').find({}).toArray();
    console.log(`✅ Created products: ${productsResult.insertedCount}`);

    // ==================== PROMOTIONS ====================
    console.log('🎉 Creating promotions...');
    const now = new Date();
    const nextWeek = new Date(now.getTime() + 7 * 24 * 60 * 60 * 1000);
    const nextMonth = new Date(now.getTime() + 30 * 24 * 60 * 60 * 1000);
    
    const promotionsResult = await db.collection('promotions').insertMany([
      {
        title: '🍔 Classic Combo',
        description: 'Burger + french fries + drink with 15% discount',
        startsAt: now,
        endsAt: nextMonth,
        bannerImage: 'https://images.unsplash.com/photo-1572802419224-296b0aeee0d9?w=800',
        isActive: true,
        appliesTo: products.filter(p => 
          ['classic-burger', 'french-fries', 'coca-cola'].includes(p.slug)
        ).map(p => p._id)
      },
      {
        title: '🥤 Happy Hours',
        description: 'All drinks -20% from 14:00 to 16:00',
        startsAt: now,
        endsAt: nextWeek,
        bannerImage: 'https://images.unsplash.com/photo-1437418747212-8d9709afab22?w=800',
        isActive: true,
        appliesTo: products.filter(p => p.categoryId.equals(drinks._id)).map(p => p._id)
      },
      {
        title: '🍰 Sweet Friday',
        description: 'Every dessert for $2.50 on Friday!',
        startsAt: now,
        endsAt: nextMonth,
        bannerImage: 'https://images.unsplash.com/photo-1551024506-0bccd828d307?w=800',
        isActive: true,
        appliesTo: products.filter(p => p.categoryId.equals(desserts._id)).map(p => p._id)
      }
    ]);
    console.log(`✅ Created promotions: ${promotionsResult.insertedCount}`);

    // ==================== MOOD QUESTIONS ====================
    console.log('😊 Creating mood questions for mood quiz...');
    const moodQuestionsResult = await db.collection('mood_questions').insertMany([
      {
        order: 1,
        text: 'How do you feel today?',
        type: 'single',
        options: [
          { value: 'energetic', label: '⚡ Full of energy', tags: ['energy', 'active'] },
          { value: 'tired', label: '😴 Tired', tags: ['comfort', 'light'] },
          { value: 'happy', label: '😊 Happy', tags: ['popular', 'classic'] },
          { value: 'stressed', label: '😰 Stressed', tags: ['comfort', 'sweet'] }
        ],
        category: 'mood',
        isActive: true
      },
      {
        order: 2,
        text: 'What do you prefer?',
        type: 'single',
        options: [
          { value: 'savory', label: '🧂 Savory', tags: ['savory'] },
          { value: 'sweet', label: '🍰 Sweet', tags: ['sweet', 'dessert'] },
          { value: 'both', label: '🤷 Not sure', tags: ['popular'] }
        ],
        category: 'preference',
        isActive: true
      },
      {
        order: 3,
        text: 'How hungry are you?',
        type: 'single',
        options: [
          { value: 'very', label: '🍔🍔🍔 Very!', tags: ['comfort-food', 'hearty'] },
          { value: 'moderate', label: '🍔 Moderately', tags: ['popular'] },
          { value: 'light', label: '🥗 Want something light', tags: ['healthy', 'light'] }
        ],
        category: 'preference',
        isActive: true
      },
      {
        order: 4,
        text: 'What do you want to try?',
        type: 'multiple',
        options: [
          { value: 'burger', label: '🍔 Burger', tags: ['burgers'] },
          { value: 'drink', label: '🥤 Drink', tags: ['drinks'] },
          { value: 'dessert', label: '🍰 Dessert', tags: ['dessert', 'sweet'] },
          { value: 'snack', label: '🍟 Snack', tags: ['snacks'] }
        ],
        category: 'preference',
        isActive: true
      },
      {
        order: 5,
        text: 'What flavors do you like?',
        type: 'multiple',
        options: [
          { value: 'classic', label: '👌 Classic', tags: ['classic', 'popular'] },
          { value: 'spicy', label: '🌶️ Spicy', tags: ['spicy'] },
          { value: 'fresh', label: '🌿 Fresh', tags: ['healthy', 'fresh'] }
        ],
        category: 'preference',
        isActive: true
      },
      {
        order: 6,
        text: 'Do you follow a diet?',
        type: 'single',
        options: [
          { value: 'yes', label: '✅ Yes, I try to eat healthy food', tags: ['healthy', 'light'] },
          { value: 'no', label: '❌ No, today I want to treat myself', tags: ['comfort-food', 'popular'] },
          { value: 'sometimes', label: '🤔 Sometimes', tags: ['popular'] }
        ],
        category: 'preference',
        isActive: true
      },
      {
        order: 7,
        text: 'With whom are you today?',
        type: 'single',
        options: [
          { value: 'alone', label: '🧍 Alone', tags: ['light', 'classic'] },
          { value: 'friends', label: '👥 With friends', tags: ['popular', 'sharing'] },
          { value: 'family', label: '👨‍👩‍👧‍👦 With family', tags: ['comfort-food', 'popular'] }
        ],
        category: 'occasion',
        isActive: true
      },
      {
        order: 8,
        text: 'What is your mood?',
        type: 'single',
        options: [
          { value: 'adventurous', label: '🎢 Ready for adventure', tags: ['spicy', 'unique'] },
          { value: 'comfort', label: '🛋️ Want comfort', tags: ['comfort-food', 'classic'] },
          { value: 'celebration', label: '🎉 Celebrating!', tags: ['popular', 'special'] }
        ],
        category: 'mood',
        isActive: true
      }
    ]);
    console.log(`✅ Created ${moodQuestionsResult.insertedCount} mood questions (AI will analyze answers and recommend products)`);

    // ==================== INDEXES ====================
    console.log('📊 Creating indexes...');
    
    await db.collection('users').createIndexes([
      { key: { email: 1 }, unique: true },
      { key: { role: 1 } }
    ]);
    
    await db.collection('categories').createIndexes([
      { key: { slug: 1 }, unique: true },
      { key: { isActive: 1 } }
    ]);
    
    await db.collection('products').createIndexes([
      { key: { categoryId: 1 } },
      { key: { slug: 1 }, unique: true },
      { key: { isActive: 1 } },
      { key: { tags: 1 } }
    ]);
    
    await db.collection('promotions').createIndexes([
      { key: { isActive: 1 } },
      { key: { startsAt: 1, endsAt: 1 } }
    ]);
    
    await db.collection('carts').createIndexes([
      { key: { userId: 1 } },
      { key: { sessionId: 1 } }
    ]);
    
    await db.collection('orders').createIndexes([
      { key: { userId: 1 } },
      { key: { status: 1 } },
      { key: { createdAt: -1 } }
    ]);
    
    await db.collection('mood_questions').createIndexes([
      { key: { order: 1 } },
      { key: { isActive: 1 } }
    ]);
    
    await db.collection('mood_rules').createIndexes([
      { key: { priority: -1 } },
      { key: { isActive: 1 } }
    ]);
    
    await db.collection('ai_sessions').createIndexes([
      { key: { userId: 1 } },
      { key: { sessionId: 1 } },
      { key: { createdAt: -1 } }
    ]);
    
    console.log('✅ Indexes created');

    // ==================== SUMMARY ====================
    console.log('\n✨ Seeding completed successfully! ✨\n');
    console.log('📊 Statistics:');
    console.log(`   👥 Users: ${await db.collection('users').countDocuments()}`);
    console.log(`   📂 Categories: ${await db.collection('categories').countDocuments()}`);
    console.log(`   🍔 Products: ${await db.collection('products').countDocuments()}`);
    console.log(`   🎉 Promotions: ${await db.collection('promotions').countDocuments()}`);
    console.log(`   😊 Questions: ${await db.collection('mood_questions').countDocuments()}`);
    console.log(`   📋 Rules: ${await db.collection('mood_rules').countDocuments()}`);
    console.log('\n🔑 Admin credentials:');
    console.log('   Email: admin@local');
    console.log('   Password: Admin123!');
    console.log('');
    
  } catch (error) {
    console.error('❌ Error during seeding:', error);
    process.exit(1);
  } finally {
    await client.close();
    console.log('🔌 Connection closed');
  }
}

// Run
seed();

