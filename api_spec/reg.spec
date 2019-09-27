// -- Requests

POST http://demo6.alpha.vkhackathon.com:8844/reg/form
Request -> 
    object: RegistrForm
<- Response
    empty

POST http://demo6.alpha.vkhackathon.com:8844/reg/status
Request -> 
    object: RegistrForm
<- Response
    empty

// -- Models

struct RegistrForm {

    userInfo: struct {
        // Имя
        firstName: String
        // Фамилия
        secondName: String
        // Отчество
        lastName: String
        // Адрес регистрации    
        registerAdress: String
        // Адрес фактического проживания
        // Если nil, то совпадает с `registerAdress`
        livingAddress: String?
        // Контактный номер телефона
        phone: String
        // Контакный email
        email: String
        // Дата рождения
        birthday: Date
        // Вероисповедание
        religion: String
        // Является гражданином страны 
        isRFCitizen: Bool
    }
    
    healtInfo: struct {
        // Статус здоровья
        healtStatus: HealthStatus
        // Болезни
        desease: String
    }

    educationInfo: struct {
        // Тип образования
        educationType: EducationType
        // Название заведения
        orgName: String
        // Полученная специальность, степени или другие квалификации
        degree: String
    }

    jobInfo: struct {
        // Название организации
        orgName: String
        // Контакты
        contacts: String
        // Должность
        position: String
        // Обязанности
        responsibilities: String
        // Текущий график работы
        workTimeTable: String
    }
    
    hobbyInfo: String

    familyInfo: struct {
        // Семейное положение
        status: FamilyStatus
        // Имя партнера
        partnerName: String?
        // ПОл партнера
        partnerSex: Sex?
        // Возраст партнера
        partnerOld: Int?
        // Кем приходится
        relationships: String?
    }

    workWithChilredExpInfo: struct? {
        // Имя организации
        orgName: String?
        // Контактные данные
        contacts: String?
        // Должность
        position: String?
        // Обязанности
        responsibilities: String
        // Возрастной период детей
        childOldGroup String
    }

    programInfo: struct {
        type: MentorType
        reason: String
        recomendation: String
        selfRecomendation: String
        whyYouWantToHelp: String
        childType: struct {
            old: Int
            sex: Sex?
            requirements: String
            visitsFrequency: String
            ifChildIsBroken: Bool
        }
    }

    lawsInfo: struct {
        useAlcohol: String
        usePsychotropic: String
        useCigarets: String
        useDrags: String
        haveACrimeRecords: String
        haveParentalRights: String
        applyReportsRight: Bool
        applyNoPrivacyRight: Bool
        whereYouFindInfo: String
    }
}

// Статус здоровья
enum HealthStatus {
    // Отлично
    case great = 0
    // Хорошо
    case good = 1
    // Средне
    case middle = 2
    // Плохо
    case bad = 3
}

// Тип образования
enum EducationType {
    // Только школа
    case school = 0
    // Универ, колледж, технарь
    case university = 1
    // Доп курсы и проч
    case specialCourses = 2
}

enum FamilyStatus {
    // В браке
    case married = 0
    // Гражд. брак
    case civilMarriage = 1
    // Разведен
    case divorced = 2
    // Вдова
    case dowager = 3
    // Не женат
    case notMarried = 4
}

enum Sex {
    case male = 0
    case female = 1
}

enum MentorType {
    case mentor = 0
    case volunteer = 1
    case partner = 2
}

struct CheckStatusRequest {
    // email пользователя
    value: String
}

struct CheckStatusResponse {
    // Статус формы пользователя
    status: FormStatus
}

enum FormStatus {
    case new = 0
    case viewed = 1
    case called = 2
    case approved = 3
}