# THIS PROJECT IS ARCHIVED AND THE API IS CURRENTLY DOWN

# PSE STOCKS API
PSE Stocks API provides access to up-to-date stock prices of all the publicly listed companies in the Philippine Stock Exchange. It contains the historical data of the market as early as December 2011. All prices are updated daily after the PSE market closes at around 3:30 PM GMT +8 Philippine Time.

___
### Support the project
**Looking for a real-time stock prices?** Currently, the API only provides the daily opening and closing prices of each stock. Hopefully, I can develop it further to support real-time stock prices updates. However, maintaining a virtual machine and cloud storage can be costly thus I would need your support in order to make it possible.

**Fund the project by sending any amount on my [Paypal Account](https://www.paypal.com/paypalme/shanemaglangit) and get notified about feature updates**

---
### Usage Guides

1. **[Get current stock price](#current-stock-price)**
2. **[Limit filter](#limit-filter)**
3. **[Date range filter](#date-range-filter)**

## Current Stock Price

**Syntax:** http://psestocks-api.shanemaglangit.com/{code}

*Replace code with your stock's symbol or code (e.g JFC, ALI, BDO)*
```http request
http://psestocks-api.shanemaglangit.com/JFC
```

## Limit Filter
**Limit** can be used to retrieve the past n days of historical data. Additionally, you can set it to **all** to get the complete historical data.

**Syntax:** http://psestocks-api.shanemaglangit.com/{code}?limit={limit}

*Replace limit with a number or `all`*
```http request
// GET historical data of JFC from the past 5 market days
http://psestocks-api.shanemaglangit.com/JFC?limit=5

// GET historical data of JFC since it was first listed
http://psestocks-api.shanemaglangit.com/JFC?limit=all
```

## Date range filter
You may also specify the date range of the historical data that you want.

**Syntax:** http://psestocks-api.shanemaglangit.com/{code}?start={start_date}&end={end_date}

**Date Format:** YYYY-mm-ddd<br/>
**Example:** 2021-02-05 (Feb 5, 2021)

*By default, start date is set to August 8, 1927 and end date is set to the current date*
```http request
http://psestocks-api.shanemaglangit.com/JFC?start=2021-02-05&end=2021-03-01
```
