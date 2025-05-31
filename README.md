# `CirQL`

[![Code Quality](https://img.shields.io/badge/Code_Quality-A+-brightgreen?style=for-the-badge&logo=go&logoColor=white)](https://goreportcard.com/report/github.com/lucasepe/drop)


> A simple, privacy-first command-line tool for managing contacts locally.

Ideal for users who value full control over their data while retaining the freedom to choose whether or not to use cloud services.

This CLI tool allows you to manage contact information entirely offline using a lightweight SQLite database.

- contacts can belong to one or more categories, making organization flexible and efficient 
- full-text search (FTS) is supported for fast, powerful lookups across contact data
- search for contacts with upcoming birthdays within a specified days range
- import contacts from vCard files, and export one, many, or all contacts as vCards

## ✨ Design Philosophy

This tool deliberately limits each contact to one email address, one phone number, and one physical address.

While many contact managers allow multiple entries per field, in real-world use, a single "main" contact 
point is usually what people rely on. This streamlined approach keeps the data model simple and encourages clarity.

To represent multiple roles or contexts (e.g., home vs. work), you can create separate contacts and use 
categories to distinguish them.

This "old-school" contact philosophy ensures ease of use, clean data, and minimal clutter—ideal 
for users who value simplicity and control over exhaustive detail.

## 👍 Support

All tools are completely free to use, with every feature fully unlocked and accessible.

If you find one or more of these tool helpful, please consider supporting its development with a donation.

Your contribution, no matter the amount, helps cover the time and effort dedicated to creating and maintaining these tools, ensuring they remain free and receive continuous improvements.

Every bit of support makes a meaningful difference and allows me to focus on building more tools that solve real-world challenges.

Thank you for your generosity and for being part of this journey!


[![Donate with PayPal](https://img.shields.io/badge/💸-Tip%20me%20on%20PayPal-0070ba?style=for-the-badge&logo=paypal&logoColor=white)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=FV575PVWGXZBY&source=url)

