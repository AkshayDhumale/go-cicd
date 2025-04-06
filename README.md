Sure! Here's a clean and well-structured **GitHub Wiki-style Markdown document** based on your content, with clickable links and icons for better visual appeal:

---

# 🔐 Azure Key Vault: Secrets Management Guide

This document covers how to create an Azure Key Vault, manage access (users & apps), add secrets, and enforce policies using **Azure Policy**.

---

## 📘 Table of Contents

- [🔧 Create Azure Key Vault](#-create-azure-key-vault)
- [🔑 Grant Access to Users and Applications](#-grant-access-to-users-and-applications)
  - [🛡️ Using Access Policies](#️-using-access-policies)
  - [👥 Using Azure RBAC](#-using-azure-rbac)
- [📥 Add Secrets to Key Vault](#-add-secrets-to-key-vault)
- [📏 Policy Enforcement with Azure Policy](#-policy-enforcement-with-azure-policy)

---

## 🔧 Create Azure Key Vault

1. Go to the [Azure Portal](https://portal.azure.com/).
2. Click on ➕ **Create a resource**.
3. Search for **Key Vault** and select it.
4. Fill in the details:
   - 🔤 **Name**: Unique vault name.
   - 🧾 **Subscription**: Select your subscription.
   - 📁 **Resource Group**: Choose or create a resource group.
   - 🌍 **Region**: Select your nearest region.
5. Click **Review + create**, then click **Create**.

🔗 [Official Docs: Quick Create - Portal](https://learn.microsoft.com/en-us/azure/key-vault/general/quick-create-portal)

---

## 🔑 Grant Access to Users and Applications

Access can be managed via **Access Policies** or **Azure RBAC**.

### 🛡️ Using Access Policies

1. Go to your Key Vault in the portal.
2. Click **Access Policies** → ➕ **Add Access Policy**.
3. Select:
   - 🔐 **Key/Secret/Certificate Permissions**
4. Click **Select principal** → Choose the user or app.
5. Click **Add**, then **Save**.

🔗 [Docs: Assign Access Policy](https://learn.microsoft.com/en-us/azure/key-vault/general/assign-access-policy)

---

### 👥 Using Azure RBAC

1. Open your Key Vault → Go to **Access Control (IAM)**.
2. Click ➕ **Add role assignment**.
3. Choose a role like:
   - `Key Vault Reader`
   - `Key Vault Secrets Officer`
4. Assign to **User / Group / Service Principal**.
5. Click **Review + assign**.

🔗 [Docs: RBAC Guide](https://learn.microsoft.com/en-us/azure/key-vault/general/rbac-guide)

---

## 📥 Add Secrets to Key Vault

1. Go to your Key Vault → **Secrets** → ➕ **Generate/Import**.
2. Choose **Upload Options**: `Manual`.
3. Enter:
   - 🏷️ **Name**
   - 🔐 **Value**
4. Click **Create**.

🔗 [Docs: Add Secret - Portal](https://learn.microsoft.com/en-us/azure/key-vault/secrets/quick-create-portal)

---

## 📏 Policy Enforcement with Azure Policy

Use [Azure Policy](https://portal.azure.com/#blade/Microsoft_Azure_Policy/PolicyMenuBlade/Overview) to enforce configurations and standards:

🔍 **Use cases**:
- Audit Key Vault configurations.
- Enforce access/networking/security rules.

### 🪄 Steps:

1. Go to **Azure Policy** in portal.
2. Browse for built-in Key Vault policy definitions.
3. Click **Assign** → Choose scope (subscription/resource group).

🔗 [Docs: Azure Policy + Key Vault](https://learn.microsoft.com/en-us/azure/key-vault/general/azure-policy)

---

## ✅ Summary

Azure Key Vault helps you securely store and manage sensitive data. Combine it with proper access control and policy enforcement for enterprise-grade security.

---

Let me know if you'd like this as a downloadable `.md` file or want to embed sample screenshots!
