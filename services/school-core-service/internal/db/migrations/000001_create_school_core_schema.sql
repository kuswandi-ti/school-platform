-- +goose Up
CREATE TABLE foundations (
    id UUID PRIMARY KEY,
    foundation_code VARCHAR(50) NOT NULL,
    name VARCHAR(150) NOT NULL,
    legal_name VARCHAR(200),
    address TEXT,
    phone VARCHAR(50),
    email VARCHAR(255),
    logo_file_id UUID,
    timezone VARCHAR(100) NOT NULL DEFAULT 'Asia/Jakarta',
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT foundations_foundation_code_unique UNIQUE (foundation_code),
    CONSTRAINT foundations_status_check CHECK (status IN ('active', 'inactive'))
);

CREATE INDEX foundations_status_idx ON foundations (status);

CREATE TABLE schools (
    id UUID PRIMARY KEY,
    foundation_id UUID NOT NULL REFERENCES foundations (id) ON DELETE RESTRICT,
    school_code VARCHAR(50) NOT NULL,
    name VARCHAR(150) NOT NULL,
    school_level VARCHAR(50) NOT NULL,
    npsn VARCHAR(50),
    address TEXT,
    phone VARCHAR(50),
    email VARCHAR(255),
    logo_file_id UUID,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT schools_foundation_school_code_unique UNIQUE (foundation_id, school_code),
    CONSTRAINT schools_id_foundation_unique UNIQUE (id, foundation_id),
    CONSTRAINT schools_school_level_check CHECK (
        school_level IN ('kindergarten', 'elementary', 'junior_high', 'senior_high')
    ),
    CONSTRAINT schools_status_check CHECK (status IN ('active', 'inactive'))
);

CREATE INDEX schools_foundation_id_idx ON schools (foundation_id);
CREATE INDEX schools_school_code_idx ON schools (school_code);
CREATE INDEX schools_status_idx ON schools (foundation_id, status);
CREATE UNIQUE INDEX schools_npsn_unique_idx ON schools (npsn) WHERE npsn IS NOT NULL;

CREATE TABLE academic_years (
    id UUID PRIMARY KEY,
    foundation_id UUID NOT NULL REFERENCES foundations (id) ON DELETE RESTRICT,
    name VARCHAR(50) NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    status VARCHAR(50) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT academic_years_foundation_name_unique UNIQUE (foundation_id, name),
    CONSTRAINT academic_years_id_foundation_unique UNIQUE (id, foundation_id),
    CONSTRAINT academic_years_date_range_check CHECK (end_date >= start_date),
    CONSTRAINT academic_years_status_check CHECK (status IN ('draft', 'active', 'closed', 'archived')),
    CONSTRAINT academic_years_active_status_check CHECK (NOT is_active OR status = 'active')
);

CREATE UNIQUE INDEX academic_years_one_active_per_foundation_idx
    ON academic_years (foundation_id) WHERE is_active;
CREATE INDEX academic_years_foundation_status_idx ON academic_years (foundation_id, status);

CREATE TABLE semesters (
    id UUID PRIMARY KEY,
    foundation_id UUID NOT NULL REFERENCES foundations (id) ON DELETE RESTRICT,
    academic_year_id UUID NOT NULL,
    name VARCHAR(50) NOT NULL,
    sequence INT NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    status VARCHAR(50) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT semesters_academic_year_fk FOREIGN KEY (academic_year_id, foundation_id)
        REFERENCES academic_years (id, foundation_id) ON DELETE RESTRICT,
    CONSTRAINT semesters_academic_year_sequence_unique UNIQUE (academic_year_id, sequence),
    CONSTRAINT semesters_id_scope_unique UNIQUE (id, foundation_id, academic_year_id),
    CONSTRAINT semesters_sequence_check CHECK (sequence > 0),
    CONSTRAINT semesters_date_range_check CHECK (end_date >= start_date),
    CONSTRAINT semesters_status_check CHECK (status IN ('draft', 'active', 'closed', 'archived')),
    CONSTRAINT semesters_active_status_check CHECK (NOT is_active OR status = 'active')
);

CREATE UNIQUE INDEX semesters_one_active_per_foundation_idx
    ON semesters (foundation_id) WHERE is_active;
CREATE INDEX semesters_academic_year_id_idx ON semesters (academic_year_id);
CREATE INDEX semesters_foundation_status_idx ON semesters (foundation_id, status);

CREATE TABLE students (
    id UUID PRIMARY KEY,
    foundation_id UUID NOT NULL REFERENCES foundations (id) ON DELETE RESTRICT,
    school_id UUID NOT NULL,
    student_number VARCHAR(100),
    nisn VARCHAR(100),
    full_name VARCHAR(150) NOT NULL,
    gender VARCHAR(20) NOT NULL,
    birth_place VARCHAR(100),
    birth_date DATE,
    religion VARCHAR(50),
    address TEXT,
    phone VARCHAR(50),
    email VARCHAR(255),
    photo_file_id UUID,
    status VARCHAR(50) NOT NULL,
    entry_year INT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
    CONSTRAINT students_school_fk FOREIGN KEY (school_id, foundation_id)
        REFERENCES schools (id, foundation_id) ON DELETE RESTRICT,
    CONSTRAINT students_id_foundation_unique UNIQUE (id, foundation_id),
    CONSTRAINT students_id_scope_unique UNIQUE (id, foundation_id, school_id),
    CONSTRAINT students_gender_check CHECK (gender IN ('male', 'female')),
    CONSTRAINT students_status_check CHECK (
        status IN ('active', 'inactive', 'transferred', 'graduated', 'dropped_out')
    ),
    CONSTRAINT students_entry_year_check CHECK (entry_year IS NULL OR entry_year >= 1900)
);

CREATE UNIQUE INDEX students_student_number_unique_idx
    ON students (foundation_id, school_id, student_number)
    WHERE student_number IS NOT NULL AND deleted_at IS NULL;
CREATE UNIQUE INDEX students_nisn_unique_idx
    ON students (foundation_id, nisn)
    WHERE nisn IS NOT NULL AND deleted_at IS NULL;
CREATE INDEX students_foundation_school_idx ON students (foundation_id, school_id);
CREATE INDEX students_full_name_idx ON students (foundation_id, school_id, full_name);
CREATE INDEX students_student_number_idx ON students (student_number) WHERE student_number IS NOT NULL;
CREATE INDEX students_nisn_idx ON students (nisn) WHERE nisn IS NOT NULL;
CREATE INDEX students_status_idx ON students (foundation_id, school_id, status);

CREATE TABLE guardians (
    id UUID PRIMARY KEY,
    foundation_id UUID NOT NULL REFERENCES foundations (id) ON DELETE RESTRICT,
    full_name VARCHAR(150) NOT NULL,
    relationship_type VARCHAR(50) NOT NULL,
    phone VARCHAR(50),
    email VARCHAR(255),
    address TEXT,
    occupation VARCHAR(100),
    user_id UUID,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT guardians_id_foundation_unique UNIQUE (id, foundation_id),
    CONSTRAINT guardians_relationship_type_check CHECK (
        relationship_type IN ('father', 'mother', 'guardian', 'other')
    ),
    CONSTRAINT guardians_status_check CHECK (status IN ('active', 'inactive'))
);

CREATE INDEX guardians_foundation_id_idx ON guardians (foundation_id);
CREATE INDEX guardians_phone_idx ON guardians (phone) WHERE phone IS NOT NULL;
CREATE INDEX guardians_email_idx ON guardians (email) WHERE email IS NOT NULL;
CREATE INDEX guardians_full_name_idx ON guardians (foundation_id, full_name);
CREATE INDEX guardians_user_id_idx ON guardians (user_id) WHERE user_id IS NOT NULL;

CREATE TABLE student_guardians (
    id UUID PRIMARY KEY,
    foundation_id UUID NOT NULL REFERENCES foundations (id) ON DELETE RESTRICT,
    student_id UUID NOT NULL,
    guardian_id UUID NOT NULL,
    relationship_type VARCHAR(50) NOT NULL,
    is_primary BOOLEAN NOT NULL DEFAULT FALSE,
    can_login BOOLEAN NOT NULL DEFAULT TRUE,
    can_receive_notification BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT student_guardians_student_fk FOREIGN KEY (student_id, foundation_id)
        REFERENCES students (id, foundation_id) ON DELETE CASCADE,
    CONSTRAINT student_guardians_guardian_fk FOREIGN KEY (guardian_id, foundation_id)
        REFERENCES guardians (id, foundation_id) ON DELETE CASCADE,
    CONSTRAINT student_guardians_student_guardian_unique UNIQUE (student_id, guardian_id),
    CONSTRAINT student_guardians_relationship_type_check CHECK (
        relationship_type IN ('father', 'mother', 'guardian', 'other')
    )
);

CREATE UNIQUE INDEX student_guardians_one_primary_per_student_idx
    ON student_guardians (student_id) WHERE is_primary;
CREATE INDEX student_guardians_guardian_id_idx ON student_guardians (guardian_id);
CREATE INDEX student_guardians_foundation_id_idx ON student_guardians (foundation_id);

CREATE TABLE teachers (
    id UUID PRIMARY KEY,
    foundation_id UUID NOT NULL REFERENCES foundations (id) ON DELETE RESTRICT,
    school_id UUID NOT NULL,
    employee_number VARCHAR(100),
    user_id UUID,
    full_name VARCHAR(150) NOT NULL,
    gender VARCHAR(20),
    birth_place VARCHAR(100),
    birth_date DATE,
    email VARCHAR(255),
    phone VARCHAR(50),
    address TEXT,
    photo_file_id UUID,
    employment_status VARCHAR(50),
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT teachers_school_fk FOREIGN KEY (school_id, foundation_id)
        REFERENCES schools (id, foundation_id) ON DELETE RESTRICT,
    CONSTRAINT teachers_id_scope_unique UNIQUE (id, foundation_id, school_id),
    CONSTRAINT teachers_gender_check CHECK (gender IS NULL OR gender IN ('male', 'female')),
    CONSTRAINT teachers_status_check CHECK (status IN ('active', 'inactive', 'resigned'))
);

CREATE UNIQUE INDEX teachers_employee_number_unique_idx
    ON teachers (foundation_id, school_id, employee_number) WHERE employee_number IS NOT NULL;
CREATE INDEX teachers_foundation_school_idx ON teachers (foundation_id, school_id);
CREATE INDEX teachers_full_name_idx ON teachers (foundation_id, school_id, full_name);
CREATE INDEX teachers_status_idx ON teachers (foundation_id, school_id, status);
CREATE INDEX teachers_user_id_idx ON teachers (user_id) WHERE user_id IS NOT NULL;

CREATE TABLE grade_levels (
    id UUID PRIMARY KEY,
    foundation_id UUID NOT NULL REFERENCES foundations (id) ON DELETE RESTRICT,
    school_id UUID NOT NULL,
    level_code VARCHAR(50) NOT NULL,
    name VARCHAR(100) NOT NULL,
    sequence INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT grade_levels_school_fk FOREIGN KEY (school_id, foundation_id)
        REFERENCES schools (id, foundation_id) ON DELETE RESTRICT,
    CONSTRAINT grade_levels_school_level_code_unique UNIQUE (foundation_id, school_id, level_code),
    CONSTRAINT grade_levels_id_scope_unique UNIQUE (id, foundation_id, school_id),
    CONSTRAINT grade_levels_sequence_check CHECK (sequence > 0)
);

CREATE INDEX grade_levels_school_sequence_idx ON grade_levels (foundation_id, school_id, sequence);

CREATE TABLE rooms (
    id UUID PRIMARY KEY,
    foundation_id UUID NOT NULL REFERENCES foundations (id) ON DELETE RESTRICT,
    school_id UUID NOT NULL,
    room_code VARCHAR(50) NOT NULL,
    name VARCHAR(100) NOT NULL,
    capacity INT,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT rooms_school_fk FOREIGN KEY (school_id, foundation_id)
        REFERENCES schools (id, foundation_id) ON DELETE RESTRICT,
    CONSTRAINT rooms_school_room_code_unique UNIQUE (foundation_id, school_id, room_code),
    CONSTRAINT rooms_id_scope_unique UNIQUE (id, foundation_id, school_id),
    CONSTRAINT rooms_capacity_check CHECK (capacity IS NULL OR capacity >= 0),
    CONSTRAINT rooms_status_check CHECK (status IN ('active', 'inactive'))
);

CREATE INDEX rooms_school_status_idx ON rooms (foundation_id, school_id, status);

CREATE TABLE classes (
    id UUID PRIMARY KEY,
    foundation_id UUID NOT NULL REFERENCES foundations (id) ON DELETE RESTRICT,
    school_id UUID NOT NULL,
    academic_year_id UUID NOT NULL,
    grade_level_id UUID NOT NULL,
    class_code VARCHAR(50) NOT NULL,
    name VARCHAR(100) NOT NULL,
    capacity INT,
    room_id UUID,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT classes_school_fk FOREIGN KEY (school_id, foundation_id)
        REFERENCES schools (id, foundation_id) ON DELETE RESTRICT,
    CONSTRAINT classes_academic_year_fk FOREIGN KEY (academic_year_id, foundation_id)
        REFERENCES academic_years (id, foundation_id) ON DELETE RESTRICT,
    CONSTRAINT classes_grade_level_fk FOREIGN KEY (grade_level_id, foundation_id, school_id)
        REFERENCES grade_levels (id, foundation_id, school_id) ON DELETE RESTRICT,
    CONSTRAINT classes_room_fk FOREIGN KEY (room_id, foundation_id, school_id)
        REFERENCES rooms (id, foundation_id, school_id) ON DELETE RESTRICT,
    CONSTRAINT classes_scope_class_code_unique UNIQUE (
        foundation_id, school_id, academic_year_id, class_code
    ),
    CONSTRAINT classes_id_scope_unique UNIQUE (id, foundation_id, school_id, academic_year_id),
    CONSTRAINT classes_capacity_check CHECK (capacity IS NULL OR capacity >= 0),
    CONSTRAINT classes_status_check CHECK (status IN ('active', 'inactive', 'archived'))
);

CREATE INDEX classes_foundation_school_academic_year_idx
    ON classes (foundation_id, school_id, academic_year_id);
CREATE INDEX classes_grade_level_id_idx ON classes (grade_level_id);
CREATE INDEX classes_room_id_idx ON classes (room_id) WHERE room_id IS NOT NULL;
CREATE INDEX classes_status_idx ON classes (foundation_id, school_id, status);

CREATE TABLE student_class_assignments (
    id UUID PRIMARY KEY,
    foundation_id UUID NOT NULL REFERENCES foundations (id) ON DELETE RESTRICT,
    school_id UUID NOT NULL,
    academic_year_id UUID NOT NULL,
    semester_id UUID,
    student_id UUID NOT NULL,
    class_id UUID NOT NULL,
    status VARCHAR(50) NOT NULL,
    assigned_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT student_class_assignments_student_fk FOREIGN KEY (student_id, foundation_id, school_id)
        REFERENCES students (id, foundation_id, school_id) ON DELETE RESTRICT,
    CONSTRAINT student_class_assignments_class_fk FOREIGN KEY (
        class_id, foundation_id, school_id, academic_year_id
    ) REFERENCES classes (id, foundation_id, school_id, academic_year_id) ON DELETE RESTRICT,
    CONSTRAINT student_class_assignments_semester_fk FOREIGN KEY (
        semester_id, foundation_id, academic_year_id
    ) REFERENCES semesters (id, foundation_id, academic_year_id) ON DELETE RESTRICT,
    CONSTRAINT student_class_assignments_status_check CHECK (status IN ('active', 'moved', 'completed'))
);

CREATE UNIQUE INDEX student_class_assignments_one_active_idx
    ON student_class_assignments (student_id, academic_year_id) WHERE status = 'active';
CREATE INDEX student_class_assignments_student_id_idx ON student_class_assignments (student_id);
CREATE INDEX student_class_assignments_class_id_idx ON student_class_assignments (class_id);
CREATE INDEX student_class_assignments_scope_status_idx
    ON student_class_assignments (foundation_id, school_id, academic_year_id, status);

CREATE TABLE teacher_assignments (
    id UUID PRIMARY KEY,
    foundation_id UUID NOT NULL REFERENCES foundations (id) ON DELETE RESTRICT,
    school_id UUID NOT NULL,
    academic_year_id UUID NOT NULL,
    semester_id UUID NOT NULL,
    teacher_id UUID NOT NULL,
    class_id UUID,
    subject_id UUID,
    assignment_type VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT teacher_assignments_teacher_fk FOREIGN KEY (teacher_id, foundation_id, school_id)
        REFERENCES teachers (id, foundation_id, school_id) ON DELETE RESTRICT,
    CONSTRAINT teacher_assignments_semester_fk FOREIGN KEY (
        semester_id, foundation_id, academic_year_id
    ) REFERENCES semesters (id, foundation_id, academic_year_id) ON DELETE RESTRICT,
    CONSTRAINT teacher_assignments_class_fk FOREIGN KEY (
        class_id, foundation_id, school_id, academic_year_id
    ) REFERENCES classes (id, foundation_id, school_id, academic_year_id) ON DELETE RESTRICT,
    CONSTRAINT teacher_assignments_assignment_type_check CHECK (
        assignment_type IN ('subject_teacher', 'class_teacher', 'extracurricular_coach')
    ),
    CONSTRAINT teacher_assignments_status_check CHECK (status IN ('active', 'inactive'))
);

CREATE INDEX teacher_assignments_teacher_id_idx ON teacher_assignments (teacher_id);
CREATE INDEX teacher_assignments_class_id_idx ON teacher_assignments (class_id) WHERE class_id IS NOT NULL;
CREATE INDEX teacher_assignments_subject_id_idx ON teacher_assignments (subject_id) WHERE subject_id IS NOT NULL;
CREATE INDEX teacher_assignments_scope_status_idx
    ON teacher_assignments (foundation_id, school_id, academic_year_id, semester_id, status);

CREATE TABLE homeroom_assignments (
    id UUID PRIMARY KEY,
    foundation_id UUID NOT NULL REFERENCES foundations (id) ON DELETE RESTRICT,
    school_id UUID NOT NULL,
    academic_year_id UUID NOT NULL,
    semester_id UUID,
    teacher_id UUID NOT NULL,
    class_id UUID NOT NULL,
    status VARCHAR(50) NOT NULL,
    approved_by UUID,
    approved_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT homeroom_assignments_teacher_fk FOREIGN KEY (teacher_id, foundation_id, school_id)
        REFERENCES teachers (id, foundation_id, school_id) ON DELETE RESTRICT,
    CONSTRAINT homeroom_assignments_class_fk FOREIGN KEY (
        class_id, foundation_id, school_id, academic_year_id
    ) REFERENCES classes (id, foundation_id, school_id, academic_year_id) ON DELETE RESTRICT,
    CONSTRAINT homeroom_assignments_semester_fk FOREIGN KEY (
        semester_id, foundation_id, academic_year_id
    ) REFERENCES semesters (id, foundation_id, academic_year_id) ON DELETE RESTRICT,
    CONSTRAINT homeroom_assignments_status_check CHECK (status IN ('active', 'inactive')),
    CONSTRAINT homeroom_assignments_approval_check CHECK (
        (approved_by IS NULL AND approved_at IS NULL) OR
        (approved_by IS NOT NULL AND approved_at IS NOT NULL)
    )
);

CREATE UNIQUE INDEX homeroom_assignments_one_active_per_class_idx
    ON homeroom_assignments (class_id, academic_year_id, COALESCE(semester_id, '00000000-0000-0000-0000-000000000000'::UUID))
    WHERE status = 'active';
CREATE INDEX homeroom_assignments_teacher_id_idx ON homeroom_assignments (teacher_id);
CREATE INDEX homeroom_assignments_scope_status_idx
    ON homeroom_assignments (foundation_id, school_id, academic_year_id, status);

CREATE TABLE audit_logs (
    id UUID PRIMARY KEY,
    foundation_id UUID NOT NULL REFERENCES foundations (id) ON DELETE RESTRICT,
    school_id UUID,
    actor_user_id UUID,
    actor_role VARCHAR(100),
    action VARCHAR(150) NOT NULL,
    module VARCHAR(100) NOT NULL,
    entity_type VARCHAR(100) NOT NULL,
    entity_id UUID NOT NULL,
    old_values_json JSONB,
    new_values_json JSONB,
    metadata_json JSONB,
    ip_address VARCHAR(100),
    user_agent TEXT,
    request_id VARCHAR(100) NOT NULL,
    correlation_id VARCHAR(100) NOT NULL,
    occurred_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT audit_logs_school_fk FOREIGN KEY (school_id, foundation_id)
        REFERENCES schools (id, foundation_id) ON DELETE RESTRICT
);

CREATE INDEX audit_logs_foundation_id_idx ON audit_logs (foundation_id);
CREATE INDEX audit_logs_school_id_idx ON audit_logs (school_id) WHERE school_id IS NOT NULL;
CREATE INDEX audit_logs_actor_user_id_idx ON audit_logs (actor_user_id) WHERE actor_user_id IS NOT NULL;
CREATE INDEX audit_logs_action_idx ON audit_logs (action);
CREATE INDEX audit_logs_entity_idx ON audit_logs (entity_type, entity_id);
CREATE INDEX audit_logs_occurred_at_idx ON audit_logs (occurred_at);

-- +goose Down
DROP TABLE IF EXISTS audit_logs;
DROP TABLE IF EXISTS homeroom_assignments;
DROP TABLE IF EXISTS teacher_assignments;
DROP TABLE IF EXISTS student_class_assignments;
DROP TABLE IF EXISTS classes;
DROP TABLE IF EXISTS rooms;
DROP TABLE IF EXISTS grade_levels;
DROP TABLE IF EXISTS teachers;
DROP TABLE IF EXISTS student_guardians;
DROP TABLE IF EXISTS guardians;
DROP TABLE IF EXISTS students;
DROP TABLE IF EXISTS semesters;
DROP TABLE IF EXISTS academic_years;
DROP TABLE IF EXISTS schools;
DROP TABLE IF EXISTS foundations;
