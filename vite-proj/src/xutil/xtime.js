export function GetZhWeekDay(t){
    var str =''
    var week = new Date(t).getDay();  
    switch (week) {  
        case 0 :  
                str += "日";  
                break;  
        case 1 :  
                str += "一";  
                break;  
        case 2 :  
                str += "二";  
                break;  
        case 3 :  
                str += "三";  
                break;  
        case 4 :  
                str += "四";  
                break;  
        case 5 :  
                str += "五";  
                break;  
        case 6 :  
                str += "六";  
                break;  
    }
    return str
}