-- +goose Up
INSERT INTO roles (id, code, name, description, is_system)
VALUES
    ('10000000-0000-0000-0000-000000000001', 'admin_yayasan', 'Admin Yayasan', 'Foundation-wide platform administration.', TRUE),
    ('10000000-0000-0000-0000-000000000002', 'kepala_sekolah', 'Kepala Sekolah', 'School leadership and approval access.', TRUE),
    ('10000000-0000-0000-0000-000000000003', 'tu_staff', 'TU/Staff', 'School administration access.', TRUE),
    ('10000000-0000-0000-0000-000000000004', 'bendahara_sekolah', 'Bendahara Sekolah', 'School finance operations access.', TRUE),
    ('10000000-0000-0000-0000-000000000005', 'guru', 'Guru', 'Assigned teaching and academic access.', TRUE),
    ('10000000-0000-0000-0000-000000000006', 'orang_tua', 'Orang Tua', 'Access to assigned student information.', TRUE),
    ('10000000-0000-0000-0000-000000000007', 'siswa', 'Siswa', 'Access to the assigned student profile.', TRUE)
ON CONFLICT (code) DO UPDATE SET
    name = EXCLUDED.name,
    description = EXCLUDED.description,
    is_system = TRUE,
    updated_at = CURRENT_TIMESTAMP;

INSERT INTO permissions (id, code, name, description, module)
VALUES
    ('20000000-0000-0000-0000-000000000001', 'identity.context.view', 'View Identity Context', 'View the authenticated user context.', 'identity'),
    ('20000000-0000-0000-0000-000000000002', 'identity.role.view', 'View Roles', 'View role and permission definitions.', 'identity'),
    ('20000000-0000-0000-0000-000000000003', 'identity.role.assign', 'Assign Roles', 'Assign scoped roles to users.', 'identity'),
    ('20000000-0000-0000-0000-000000000004', 'school.foundation.view', 'View Foundation', 'View foundation information.', 'school'),
    ('20000000-0000-0000-0000-000000000005', 'school.foundation.manage', 'Manage Foundation', 'Manage foundation information.', 'school'),
    ('20000000-0000-0000-0000-000000000006', 'school.school.view', 'View Schools', 'View school information.', 'school'),
    ('20000000-0000-0000-0000-000000000007', 'school.school.manage', 'Manage Schools', 'Manage school information.', 'school'),
    ('20000000-0000-0000-0000-000000000008', 'school.student.view', 'View Students', 'View students within assigned scope.', 'school'),
    ('20000000-0000-0000-0000-000000000009', 'school.student.manage', 'Manage Students', 'Manage students within assigned scope.', 'school'),
    ('20000000-0000-0000-0000-000000000010', 'school.teacher.view', 'View Teachers', 'View teachers within assigned scope.', 'school'),
    ('20000000-0000-0000-0000-000000000011', 'school.teacher.manage', 'Manage Teachers', 'Manage teachers within assigned scope.', 'school'),
    ('20000000-0000-0000-0000-000000000012', 'school.class.view', 'View Classes', 'View classes within assigned scope.', 'school'),
    ('20000000-0000-0000-0000-000000000013', 'school.class.manage', 'Manage Classes', 'Manage classes within assigned scope.', 'school'),
    ('20000000-0000-0000-0000-000000000014', 'admission.applicant.view', 'View Applicants', 'View applicants within assigned scope.', 'admission'),
    ('20000000-0000-0000-0000-000000000015', 'admission.applicant.manage', 'Manage Applicants', 'Manage applicants within assigned scope.', 'admission'),
    ('20000000-0000-0000-0000-000000000016', 'admission.applicant.decide', 'Decide Applicants', 'Accept or reject applicants.', 'admission'),
    ('20000000-0000-0000-0000-000000000017', 'finance.bill.view', 'View Bills', 'View bills within assigned scope.', 'finance'),
    ('20000000-0000-0000-0000-000000000018', 'finance.bill.manage', 'Manage Bills', 'Manage bills within assigned scope.', 'finance'),
    ('20000000-0000-0000-0000-000000000019', 'finance.payment.view', 'View Payments', 'View payments within assigned scope.', 'finance'),
    ('20000000-0000-0000-0000-000000000020', 'finance.payment.verify', 'Verify Payments', 'Verify or reject payment submissions.', 'finance'),
    ('20000000-0000-0000-0000-000000000021', 'academic.schedule.view', 'View Schedules', 'View schedules within assigned scope.', 'academic'),
    ('20000000-0000-0000-0000-000000000022', 'academic.schedule.manage', 'Manage Schedules', 'Manage school schedules.', 'academic'),
    ('20000000-0000-0000-0000-000000000023', 'academic.attendance.view', 'View Attendance', 'View attendance within assigned scope.', 'academic'),
    ('20000000-0000-0000-0000-000000000024', 'academic.attendance.manage', 'Manage Attendance', 'Record attendance within assigned scope.', 'academic'),
    ('20000000-0000-0000-0000-000000000025', 'academic.grade.view', 'View Grades', 'View grades within assigned scope.', 'academic'),
    ('20000000-0000-0000-0000-000000000026', 'academic.grade.manage', 'Manage Grades', 'Record grades within assigned scope.', 'academic'),
    ('20000000-0000-0000-0000-000000000027', 'academic.report_card.view', 'View Report Cards', 'View report cards within assigned scope.', 'academic'),
    ('20000000-0000-0000-0000-000000000028', 'academic.report_card.publish', 'Publish Report Cards', 'Publish approved report cards.', 'academic'),
    ('20000000-0000-0000-0000-000000000029', 'communication.announcement.view', 'View Announcements', 'View announcements within assigned scope.', 'communication'),
    ('20000000-0000-0000-0000-000000000030', 'communication.announcement.manage', 'Manage Announcements', 'Create and publish announcements.', 'communication'),
    ('20000000-0000-0000-0000-000000000031', 'reporting.dashboard.view', 'View Dashboards', 'View dashboards within assigned scope.', 'reporting')
ON CONFLICT (code) DO UPDATE SET
    name = EXCLUDED.name,
    description = EXCLUDED.description,
    module = EXCLUDED.module,
    updated_at = CURRENT_TIMESTAMP;

WITH role_permission_codes (role_code, permission_code) AS (
    SELECT 'admin_yayasan', code
    FROM permissions
    UNION ALL VALUES
    ('kepala_sekolah', 'identity.context.view'),
    ('kepala_sekolah', 'school.foundation.view'),
    ('kepala_sekolah', 'school.school.view'),
    ('kepala_sekolah', 'school.student.view'),
    ('kepala_sekolah', 'school.teacher.view'),
    ('kepala_sekolah', 'school.class.view'),
    ('kepala_sekolah', 'admission.applicant.view'),
    ('kepala_sekolah', 'admission.applicant.decide'),
    ('kepala_sekolah', 'finance.bill.view'),
    ('kepala_sekolah', 'finance.payment.view'),
    ('kepala_sekolah', 'academic.schedule.view'),
    ('kepala_sekolah', 'academic.attendance.view'),
    ('kepala_sekolah', 'academic.grade.view'),
    ('kepala_sekolah', 'academic.report_card.view'),
    ('kepala_sekolah', 'academic.report_card.publish'),
    ('kepala_sekolah', 'communication.announcement.view'),
    ('kepala_sekolah', 'communication.announcement.manage'),
    ('kepala_sekolah', 'reporting.dashboard.view'),
    ('tu_staff', 'identity.context.view'),
    ('tu_staff', 'school.school.view'),
    ('tu_staff', 'school.student.view'),
    ('tu_staff', 'school.student.manage'),
    ('tu_staff', 'school.teacher.view'),
    ('tu_staff', 'school.teacher.manage'),
    ('tu_staff', 'school.class.view'),
    ('tu_staff', 'school.class.manage'),
    ('tu_staff', 'admission.applicant.view'),
    ('tu_staff', 'admission.applicant.manage'),
    ('tu_staff', 'academic.schedule.view'),
    ('tu_staff', 'academic.schedule.manage'),
    ('tu_staff', 'communication.announcement.view'),
    ('bendahara_sekolah', 'identity.context.view'),
    ('bendahara_sekolah', 'school.school.view'),
    ('bendahara_sekolah', 'school.student.view'),
    ('bendahara_sekolah', 'finance.bill.view'),
    ('bendahara_sekolah', 'finance.bill.manage'),
    ('bendahara_sekolah', 'finance.payment.view'),
    ('bendahara_sekolah', 'finance.payment.verify'),
    ('bendahara_sekolah', 'reporting.dashboard.view'),
    ('guru', 'identity.context.view'),
    ('guru', 'school.student.view'),
    ('guru', 'school.class.view'),
    ('guru', 'academic.schedule.view'),
    ('guru', 'academic.attendance.view'),
    ('guru', 'academic.attendance.manage'),
    ('guru', 'academic.grade.view'),
    ('guru', 'academic.grade.manage'),
    ('guru', 'academic.report_card.view'),
    ('guru', 'communication.announcement.view'),
    ('orang_tua', 'identity.context.view'),
    ('orang_tua', 'finance.bill.view'),
    ('orang_tua', 'finance.payment.view'),
    ('orang_tua', 'academic.schedule.view'),
    ('orang_tua', 'academic.attendance.view'),
    ('orang_tua', 'academic.grade.view'),
    ('orang_tua', 'academic.report_card.view'),
    ('orang_tua', 'communication.announcement.view'),
    ('siswa', 'identity.context.view'),
    ('siswa', 'academic.schedule.view'),
    ('siswa', 'academic.attendance.view'),
    ('siswa', 'academic.grade.view'),
    ('siswa', 'academic.report_card.view'),
    ('siswa', 'communication.announcement.view')
)
INSERT INTO role_permissions (id, role_id, permission_id)
SELECT
    (SUBSTR(MD5(r.code || ':' || p.code), 1, 8) || '-' ||
     SUBSTR(MD5(r.code || ':' || p.code), 9, 4) || '-' ||
     SUBSTR(MD5(r.code || ':' || p.code), 13, 4) || '-' ||
     SUBSTR(MD5(r.code || ':' || p.code), 17, 4) || '-' ||
     SUBSTR(MD5(r.code || ':' || p.code), 21, 12))::UUID,
    r.id,
    p.id
FROM role_permission_codes rpc
JOIN roles r ON r.code = rpc.role_code
JOIN permissions p ON p.code = rpc.permission_code
ON CONFLICT DO NOTHING;

-- +goose Down
DELETE FROM role_permissions
WHERE role_id IN (
    SELECT id FROM roles
    WHERE code IN ('admin_yayasan', 'kepala_sekolah', 'tu_staff', 'bendahara_sekolah', 'guru', 'orang_tua', 'siswa')
)
AND permission_id IN (
    SELECT id FROM permissions WHERE id::text LIKE '20000000-0000-0000-0000-%'
);

DELETE FROM permissions WHERE id::text LIKE '20000000-0000-0000-0000-%';
DELETE FROM roles WHERE code IN ('admin_yayasan', 'kepala_sekolah', 'tu_staff', 'bendahara_sekolah', 'guru', 'orang_tua', 'siswa');
