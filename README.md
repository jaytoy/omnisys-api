# This repo is no longer maintained

The new repository can be found at: [https://github.com/jaytoy/omni-api](https://github.com/jaytoy/omni-api)

# Introduction

This is a demo API that I built for my bachelor thesis “Omnichannel Architecture for Retail SMEs”. The thesis is trying to solve the challenges that SMEs face when selling in omnichannel. The idea is to create an API to connect to a headless e-commerce storefront and a POS system. 

The e-commerce storefront and POS system are the frontend layer of the application. The former provides the user interface for customers to view the products and place an order, while the latter provides the user interface for the store manager or sales associates to process purchases and returns. The backend layer is an API to handle all the requests from the e-commerce website and POS system.

![High-level-solution.drawio (1).png](README%20548e20714d2d449f91d7e276abf99f31/High-level-solution.drawio_(1).png)

# What is Omnichannel Retailing?

Omnichannel Retailing provides the greatest consumer experience by allowing customers, information, and inventory to move seamlessly between numerous channels throughout the purchasing process. The key point for Omnichannel Retailing is an integrated/seamless experience across all channels.

# Why Clean Architecture?

The API was designed and implemented based on [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).

## 1. Highly decoupled from details

Details are the elements that are not related to the business domain or business logic, so they belong to the outermost layer in the clean architecture. The inner layers know nothing about the details that the outer layer is implementing.

## 2. Easily testable

The business logic is separated from external frameworks or libraries, so there is no hard coupling between services, making the system testable.

## 3. Maintainable and easier to understand

In clean architecture, separation of concern is achieved as the code is separated into four layers depending on responsibilities. The dependency rule in clean architecture keeps the dependency direction from outside to inside, so these elements can be changed easily without affecting the system.
